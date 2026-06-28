package worker

import (
	inputport "emailservice/core/application/ports/input"
	"log/slog"
)

// NewSendEmailTaskHandler initializes and returns a new instance of SendEmailTaskHandler
// ready for production use.
//
// RECOMMENDATION: Always use this factory function to instantiate the handler instead of
// initializing the struct directly to ensure all core dependencies are properly wired.
func NewSendEmailTaskHandler(
	usecase inputport.ExecuteSendEmailPort,
	logger *slog.Logger,
) *SendEmailTaskHandler {
	return &SendEmailTaskHandler{
		UseCase: usecase,
		Logger:  logger,
	}
}
