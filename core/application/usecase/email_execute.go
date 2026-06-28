package usecase

import (
	"emailservice/core/application/message"
	"emailservice/core/application/ports/output"
)

// ExecuteSendEmailUsecase implements the use case responsible for executing
// the synchronous sending of an email.
//
// This use case orchestrates the email sending flow by first rendering
// the email content and then delegating the delivery to an output port.
type ExecuteSendEmailUsecase struct {
	// Sender is responsible for delivering the rendered email content
	// through an external email service.
	Sender outputport.SendEmailPort

	// Renderer is responsible for generating the email body based on the Message.
	Renderer outputport.RenderEmailContentPort
}

// Execute renders the email content and sends the email synchronously.
//
// Errors:
//   - Returns apperrors.InfrastructureError if no template is registered for the message type.
//   - Returns apperrors.InfrastructureError if the template cannot be parsed.
//   - Returns apperrors.InfrastructureError if failed to render email template.
//   - Returns apperrors.InfrastructureError. Returns wrapped with apperrors.ErrTemporaryFailure
//     on rate limit breaches, or wrapped with apperrors.ErrPermanentFailure
//     on API or network failures.
func (e *ExecuteSendEmailUsecase) Execute(message message.Message) error {
	subject, body, err := e.Renderer.Render(message)
	if err != nil {
		return err
	}

	err = e.Sender.SendEmail(message.To, subject, body)
	if err != nil {
		return err
	}

	return nil
}
