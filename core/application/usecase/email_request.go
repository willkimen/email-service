package usecase

import (
	"emailservice/core/application/message"
	"emailservice/core/application/ports/output"
)

// RequestSendEmailUseCase implements the use case responsible for requesting
// the asynchronous sending of an email.
//
// This use case and delegates the publishing of the request to an output port.
// It does not send the email directly.
type RequestSendEmailUseCase struct {
	// Publisher publishes a request to send an email, typically by
	// enqueueing a message or task for asynchronous processing.
	Publisher outputport.PublishEmailRequestPort
}

// Execute and publishes a request for it to be sent asynchronously.
//
// Errors:
//
//   - Returns apperrors.InfrastructureError if the message fails to marshal into JSON.
//   - Returns apperrors.InfrastructureError if the background worker client fails to
//     enqueue the task.
func (re *RequestSendEmailUseCase) Execute(message message.Message) error {
	// Publishing the request delegates the responsibility of delivery
	// to an asynchronous mechanism such as a queue or task scheduler.
	if err := re.Publisher.Publish(message); err != nil {
		return err
	}

	return nil
}
