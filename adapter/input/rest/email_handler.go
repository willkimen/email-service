package rest

import "emailservice/core/application/ports/input"

// HandlerEmail handles HTTP requests related to email sending operations.
//
// It acts as an input adapter, receiving HTTP requests, converting them
// into domain messages, and delegating execution to the appropriate
// application use case.
type HandlerEmail struct {
	Usecase inputport.RequestEmailSendUseCase
}

func NewHandlerEmail(usecase inputport.RequestEmailSendUseCase) *HandlerEmail {
	return &HandlerEmail{Usecase: usecase}
}
