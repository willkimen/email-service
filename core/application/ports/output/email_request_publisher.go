package outputport

import "emailservice/core/application/message"

// PublishEmailRequestPort defines the output port responsible for publishing
// email send requests to an asynchronous processing mechanism.
//
// Implementations of this interface are expected to take a validated
// Message and publish it to a queue, broker, or task system so that
// the actual email sending can be handled asynchronously.
type PublishEmailRequestPort interface {
	// Publish enqueues message for asynchronous processing.
	//
	// Errors:
	//   - Returns apperrors.InfrastructureError if the message fails to marshal into JSON.
	//   - Returns apperrors.InfrastructureError if the background worker client fails
	//     to enqueue the task.
	Publish(message message.Message) error
}
