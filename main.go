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
	"github.com/resend/resend-go/v3"
)

func rootComposition(baseLogger *slog.Logger) (
	*http.Server, *asynq.Server, *asynq.ServeMux, *asynq.Client) {
	// ===== ASYNQ CLIENT (publisher) =====
	redisOpt := asynq.RedisClientOpt{
		Addr: os.Getenv("BROKER_ADDR"),
	}

	asynqClient := asynq.NewClient(redisOpt)

	// ===== RESEND CLIENT =====
	resendClient := resend.NewClient(os.Getenv("RESEND_API_KEY"))

	// ===== OUTPUT ADAPTERS =====
	publisherOutputAdapter := emailpublisher.NewAsynqEmailPublisherAdapter(
		asynqClient,
		baseLogger,
	)
	rendererOutputAdapter := renderer.NewHTMLEmailContentRendererAdapter(baseLogger)
	senderOutputAdapter := emailsender.NewResendEmailSenderAdapter(
		resendClient,
		os.Getenv("FROM_EMAIL"),
		baseLogger,
	)
	loggerOutputAdapter := sloglogger.NewSlogLogger(baseLogger)

	// ===== USE CASES =====
	requestUsecase := usecase.NewRequestSendEmailUseCase(
		publisherOutputAdapter,
		loggerOutputAdapter,
	)
	executeUsecase := usecase.NewExecuteSendEmailUseCase(
		senderOutputAdapter,
		rendererOutputAdapter,
		loggerOutputAdapter,
	)

	// ===== INPUT ADAPTER: HTTP SERVER =====
	httpHandlerInputAdapter := rest.NewSendEmailHandler(requestUsecase, baseLogger)

	httpServer := &http.Server{
		Addr:     ":8080",
		Handler:  httpHandlerInputAdapter.Routes(),
		ErrorLog: slog.NewLogLogger(baseLogger.Handler(), slog.LevelError),
	}

	// ===== INPUT ADAPTER: ASYNQ SERVER (worker) =====
	asynqMux := asynq.NewServeMux()
	asynqHandlerInputAdapter := worker.NewSendEmailTaskHandler(
		executeUsecase,
		baseLogger,
	)

	asynqMux.HandleFunc(
		"email:send",
		asynqHandlerInputAdapter.ProcessSendEmail,
	)

	asynqServer := asynq.NewServer(
		redisOpt,
		asynq.Config{
			Concurrency: 5,
		},
	)

	return httpServer, asynqServer, asynqMux, asynqClient
}

func main() {
	baseLogger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	baseLogger.Info("starting email service")

	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer stop()

	httpServer, asynqServer, asynqMux, asynqClient := rootComposition(baseLogger)
	defer asynqClient.Close()

	httpErrCh := make(chan error, 1)
	asynqErrCh := make(chan error, 1)

	// ===== START HTTP SERVER =====
	go func() {
		baseLogger.Info("http server starting", "addr", ":8080")
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
