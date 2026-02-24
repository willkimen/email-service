package usecase

import (
	"emailservice/core/application/email_message"
	"emailservice/core/application/ports/output"
	"fmt"
)

// ExecuteSendEmailUsecase implements the use case responsible for executing
// the synchronous sending of an email.
//
// This use case orchestrates the email sending flow by first rendering
// the email content and then delegating the delivery to an output port.
type ExecuteSendEmailUsecase struct {
	// Sender is responsible for delivering the rendered email content
	// through an external email service.
	Sender outputport.SendEmailOutputPort
	// Renderer is responsible for generating the email body based on
	// the EmailMessage.
	Renderer outputport.RenderEmailContentOutputPort
}

func NewExecuteSendEmailUseCase(
	sender outputport.SendEmailOutputPort,
	renderer outputport.RenderEmailContentOutputPort,
) *ExecuteSendEmailUsecase {
	return &ExecuteSendEmailUsecase{
		Sender:   sender,
		Renderer: renderer,
	}

}

// Execute renders the email content and sends the email synchronously.
func (e *ExecuteSendEmailUsecase) ExecuteSend(message emailmessage.EmailMessage) error {
	body, err := e.Renderer.Render(message)
	if err != nil {
		return fmt.Errorf("send email failed during rendering: %w", err)
	}

	err = e.Sender.SendEmail(message.GetTo(), message.GetSubject(), body)
	if err != nil {
		return fmt.Errorf("send email failed during sending: %w", err)
	}

	return nil
}
