package rest

import "emailservice/core/application/ports/input"

// SendEmailHandler handles HTTP requests related to email sending operations.
//
// It acts as an input adapter, receiving HTTP requests, converting them
// into domain messages, and delegating execution to the appropriate
// application use case.
type SendEmailHandler struct {
	Usecase inputport.RequestSendEmailInputPort
}

func NewSendEmailHandler(usecase inputport.RequestSendEmailInputPort) *SendEmailHandler {
	return &SendEmailHandler{Usecase: usecase}
}
