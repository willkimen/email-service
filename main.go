package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"emailservice/adapter/input/rest"
	"emailservice/adapter/input/worker"
	"emailservice/adapter/output/asynq_publisher"
	"emailservice/adapter/output/content_renderer/html"
	"emailservice/adapter/output/resend"
	"emailservice/core/application/usecase"

	"github.com/hibiken/asynq"
	"github.com/joho/godotenv"
	"github.com/resend/resend-go/v3"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("could not load .env")
	}

	// contexto raiz do app
	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer stop()

	redisOpt := asynq.RedisClientOpt{
		Addr: os.Getenv("BROKER_ADDR"),
	}

	// ===== ASYNQ CLIENT (publisher) =====
	asynqClient := asynq.NewClient(redisOpt)
	defer asynqClient.Close()

	// ===== RESEND CLIENTS =====
	resendClient := resend.NewClient(os.Getenv("RESEND_API_KEY"))

	// ===== OUTPUT ADAPTERS =====
	publisherAdapter := emailpublisher.NewAsynqEmailPublisherAdapter(asynqClient)
	rendererAdapter := renderer.NewHTMLEmailContentRendererAdapter()
	senderAdapter := emailsender.NewResendEmailSenderAdapter(
		resendClient,
		os.Getenv("FROM_EMAIL"),
	)

	// ===== USE CASES =====
	requestUsecase := usecase.NewRequestSendEmailUseCase(publisherAdapter)
	executeUsecase := usecase.NewExecuteSendEmailUseCase(
		senderAdapter,
		rendererAdapter,
	)

	// ===== INPUT ADAPTER: HTTP SERVER =====
	httpHandler := rest.NewSendEmailHandler(requestUsecase)

	httpServer := &http.Server{
		Addr:    ":4000",
		Handler: httpHandler.Routes(),
	}

	// ===== INPUT ADAPTER: ASYNQ SERVER (worker) =====
	asynqMux := asynq.NewServeMux()
	asynqHandler := worker.NewSendEmailTaskHandler(executeUsecase)

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
		log.Println("HTTP server started on :4000")
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			httpErrCh <- err
		}
	}()

	// ===== START ASYNQ WORKER =====
	go func() {
		log.Println("Asynq worker started")
		if err := asynqServer.Run(asynqMux); err != nil {
			asynqErrCh <- err
		}
	}()

	// ===== WAIT =====
	select {
	case <-ctx.Done():
		log.Println("shutdown signal received")
	case err := <-httpErrCh:
		log.Printf("http server error: %v", err)

	case err := <-asynqErrCh:
		log.Printf("asynq server error: %v", err)
	}

	// ===== GRACEFUL SHUTDOWN =====
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	log.Println("shutting down HTTP server")
	if err := httpServer.Shutdown(shutdownCtx); err != nil {
		log.Printf("http shutdown error: %v", err)
	}

	log.Println("shutting down Asynq server")
	asynqServer.Shutdown()

	log.Println("application stopped")
}
