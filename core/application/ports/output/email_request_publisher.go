package outputport

import "emailservice/core/application/email_message"

// PublishEmailRequestOutputPort defines the output port responsible for publishing
// email send requests to an asynchronous processing mechanism.
//
// Implementations of this interface are expected to take a validated
// EmailMessage and publish it to a queue, broker, or task system so that
// the actual email sending can be handled asynchronously.
type PublishEmailRequestOutputPort interface {
	// Publish sends the given EmailMessage to the underlying messaging
	// or task infrastructure.
	//
	// It returns an error if the message cannot be published.
	Publish(message emailmessage.EmailMessage) error
}
