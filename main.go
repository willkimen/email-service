package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"emailservice/adapter/input/rest"
	"emailservice/adapter/input/worker"
	"emailservice/adapter/output/asynq_publisher"
	"emailservice/adapter/output/content_renderer/html"
	"emailservice/adapter/output/logger"
	"emailservice/adapter/output/resend"
	"emailservice/core/application/usecase"

	"github.com/hibiken/asynq"
	"github.com/joho/godotenv"
	"github.com/resend/resend-go/v3"
)

func main() {
	baseLogger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	baseLogger.Info("starting email service")

	if err := godotenv.Load(".env"); err != nil {
		baseLogger.Warn("could not load .env file")
	}

	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer stop()

	// ===== ASYNQ CLIENT (publisher) =====
	redisOpt := asynq.RedisClientOpt{
		Addr: os.Getenv("BROKER_ADDR"),
	}

	asynqClient := asynq.NewClient(redisOpt)
	defer asynqClient.Close()

	// ===== RESEND CLIENTS =====
	resendClient := resend.NewClient(os.Getenv("RESEND_API_KEY"))

	// ===== OUTPUT ADAPTERS =====
	publisherAdapter := emailpublisher.NewAsynqEmailPublisherAdapter(asynqClient, baseLogger)
	rendererAdapter := renderer.NewHTMLEmailContentRendererAdapter(baseLogger)
	senderAdapter := emailsender.NewResendEmailSenderAdapter(
		resendClient,
		os.Getenv("FROM_EMAIL"),
		baseLogger,
	)
	loggerAdapter := sloglogger.New(baseLogger)

	// ===== USE CASES =====
	requestUsecase := usecase.NewRequestSendEmailUseCase(
		publisherAdapter,
		loggerAdapter,
	)
	executeUsecase := usecase.NewExecuteSendEmailUseCase(
		senderAdapter,
		rendererAdapter,
		loggerAdapter,
	)

	// ===== INPUT ADAPTER: HTTP SERVER =====
	httpHandler := rest.NewSendEmailHandler(requestUsecase, baseLogger)

	httpServer := &http.Server{
		Addr:    ":4000",
		Handler: httpHandler.Routes(),
	}

	// ===== INPUT ADAPTER: ASYNQ SERVER (worker) =====
	asynqMux := asynq.NewServeMux()
	asynqHandler := worker.NewSendEmailTaskHandler(
		executeUsecase,
		baseLogger,
	)

	asynqMux.HandleFunc(
		"email:send",
		asynqHandler.ProcessSendEmail,
	)

	asynqServer := asynq.NewServer(
		redisOpt,
		asynq.Config{
			Concurrency: 5,
		},
	)

	httpErrCh := make(chan error, 1)
	asynqErrCh := make(chan error, 1)

	// ===== START HTTP SERVER =====
	go func() {
		baseLogger.Info("http server starting", "addr", ":4000")
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			httpErrCh <- err
		}
	}()

	// ===== START ASYNQ WORKER =====
	go func() {
		baseLogger.Info("asynq worker starting", "concurrency", 5)
		if err := asynqServer.Run(asynqMux); err != nil {
			asynqErrCh <- err
		}
	}()

	// ===== WAIT =====
	select {
	case <-ctx.Done():
		baseLogger.Info("shutdown signal received")

	case err := <-httpErrCh:
		baseLogger.Error("http server error", "error", err)

	case err := <-asynqErrCh:
		baseLogger.Error("asynq server error", "error", err)
	}

	// ===== GRACEFUL SHUTDOWN =====
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	baseLogger.Info("shutting down http server")
	if err := httpServer.Shutdown(shutdownCtx); err != nil {
		log.Printf("http shutdown error: %v", err)
	}

	baseLogger.Info("shutting down asynq server")
	asynqServer.Shutdown()

	baseLogger.Info("application stopped")
}
