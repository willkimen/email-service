package inputport

import "emailservice/core/application/message"

// ExecuteSendEmailPort defines the input port responsible for executing
// the email sending process.
//
// Implementations of this interface orchestrate the application flow
// for sending emails, receiving a fully constructed and valid Message,
// and delegating the operation to the appropriate output adapters.
type ExecuteSendEmailPort interface {
	// Execute triggers the email sending workflow for the given message.
	// The message must be in a valid state according to rules.
	//
	// Errors:
	//   - Returns apperrors.InfrastructureError if no template is registered
	//     for the message type.
	//   - Returns apperrors.InfrastructureError if the template cannot be parsed.
	//   - Returns apperrors.InfrastructureError if failed to render email template.
	//   - Returns apperrors.InfrastructureError. Returns wrapped with
	//     apperrors.ErrTemporaryFailure on rate limit breaches, or
	//     wrapped with apperrors.ErrPermanentFailure on API or network failures.
	Execute(message message.Message) error
}
