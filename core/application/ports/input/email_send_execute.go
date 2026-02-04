package inputport

import "emailservice/core/application/email_message"

// ExecuteSendEmailInputPort defines the input port responsible for executing
// the email sending process.
//
// Implementations of this interface orchestrate the application flow
// for sending emails, receiving a fully constructed EmailMessage,
// validating rules, and delegating the operation to the appropriate
// output adapters.
type ExecuteSendEmailInputPort interface {
	// ExecuteSend triggers the email sending workflow for the given message.
	// The message must be in a valid state according to rules.
	ExecuteSend(message emailmessage.EmailMessage) error
}
