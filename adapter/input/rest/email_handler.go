package rest

import "emailservice/core/application/ports/input"

type HandlerEmail struct {
	Usecase inputport.RequestEmailSendUseCase
}
