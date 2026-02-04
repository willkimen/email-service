package inputport

import "emailservice/core/application/email_message"

// RequestSendEmailInputPort defines the input port responsible for requesting
// the asynchronous sending of an email.
//
// Implementations of this interface validate the provided EmailMessage
// according to rules and delegate the request to an output adapter,
// such as a message queue or task publisher.
type RequestSendEmailInputPort interface {
	// Request validates the email message and requests its delivery
	// through the configured asynchronous mechanism.
	Request(message emailmessage.EmailMessage) error
}
