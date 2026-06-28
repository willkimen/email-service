package usecase

import outputport "emailservice/core/application/ports/output"

/*
NewRequestSendEmailUseCase initializes and returns a new RequestSendEmailUseCase
configured with the required outbound ports.
*/
func NewRequestSendEmailUseCase(
	publisher outputport.PublishEmailRequestPort,
) *RequestSendEmailUseCase {
	return &RequestSendEmailUseCase{Publisher: publisher}
}

/*
NewExecuteSendEmailUseCase initializes and returns a new ExecuteSendEmailUsecase
configured with the required outbound ports.
*/
func NewExecuteSendEmailUseCase(
	sender outputport.SendEmailPort,
	renderer outputport.RenderEmailContentPort,
) *ExecuteSendEmailUsecase {
	return &ExecuteSendEmailUsecase{
		Sender:   sender,
		Renderer: renderer,
	}
}
