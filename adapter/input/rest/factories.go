package rest

import (
	"emailservice/core/application/ports/input"
	"log/slog"
)

func NewSendEmailHandler(
	usecase inputport.RequestSendEmailPort,
	logger *slog.Logger,
) *SendEmailHandler {
	return &SendEmailHandler{
		RequestSendEmailUsecase: usecase,
		Logger:                  logger,
	}
}
