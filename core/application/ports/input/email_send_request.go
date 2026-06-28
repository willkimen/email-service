package inputport

import "emailservice/core/application/message"

// RequestSendEmailPort defines the input port responsible for requesting
// the asynchronous sending of an email.
//
// Implementations of this interface validate the provided message
// according to rules and delegate the request to an output adapter,
// such as a message queue or task publisher.
type RequestSendEmailPort interface {
	// Execute validates the email message and requests its delivery
	// through the configured asynchronous mechanism.
	//
	//  - Returns apperrors.InfrastructureError if the message fails to marshal into JSON.
	//  - Returns apperrors.InfrastructureError if the background worker client fails to
	//    enqueue the task.
	Execute(message message.Message) error
}
