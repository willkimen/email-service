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
	Logger   outputport.Logger
}

func NewExecuteSendEmailUseCase(
	sender outputport.SendEmailOutputPort,
	renderer outputport.RenderEmailContentOutputPort,
	logger outputport.Logger,
) *ExecuteSendEmailUsecase {
	return &ExecuteSendEmailUsecase{
		Sender:   sender,
		Renderer: renderer,
		Logger:   logger,
	}

}

// Execute renders the email content and sends the email synchronously.
func (e *ExecuteSendEmailUsecase) ExecuteSend(message emailmessage.EmailMessage) error {
	e.Logger.Info(
		"starting email send execution",
		"to", message.GetTo(),
		"subject", message.GetSubject(),
	)

	body, err := e.Renderer.Render(message)
	if err != nil {
		e.Logger.Error(
			"failed to render email content",
			err,
			"to", message.GetTo(),
			"subject", message.GetSubject(),
		)

		return fmt.Errorf("send email failed during rendering: %w", err)
	}

	err = e.Sender.SendEmail(message.GetTo(), message.GetSubject(), body)
	if err != nil {
		e.Logger.Error(
			"failed to send email",
			err,
			"to", message.GetTo(),
			"subject", message.GetSubject(),
		)
		return fmt.Errorf("send email failed during sending: %w", err)
	}

	e.Logger.Info(
		"email sent successfully",
		"to", message.GetTo(),
		"subject", message.GetSubject(),
	)

	return nil
}
