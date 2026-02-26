package rest

import (
	"emailservice/core/application/ports/input"
	"log/slog"
)

// SendEmailHandler handles HTTP requests related to email sending operations.
//
// It acts as an input adapter, receiving HTTP requests, converting them
// into messages, and delegating execution to the appropriate
// application use case.
type SendEmailHandler struct {
	Usecase inputport.RequestSendEmailInputPort
	Logger  *slog.Logger
}

func NewSendEmailHandler(
	usecase inputport.RequestSendEmailInputPort,
	logger *slog.Logger,
) *SendEmailHandler {
	return &SendEmailHandler{
		Usecase: usecase,
		Logger:  logger,
	}
}
