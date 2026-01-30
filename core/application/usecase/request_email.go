package usecase

import (
	"emailservice/core/application/email_message"
	outputport "emailservice/core/application/ports/output"
	"fmt"
)

// RequestEmail implements the use case responsible for requesting
// the asynchronous sending of an email.
//
// This use case validates the email message according to rules
// and delegates the publishing of the request to an output port.
// It does not send the email directly.
type RequestEmail struct {
	// Publisher publishes a request to send an email, typically by
	// enqueueing a message or task for asynchronous processing.
	Publisher outputport.EmailRequestPublisher
}

// Request validates the given email message and publishes a request
// for it to be sent asynchronously.
//
// The message must be in a valid state. If validation fails,
// a validation error is returned. If publishing fails, an
// infrastructure-level error is returned.
func (re *RequestEmail) Request(message emailmessage.EmailMessage) error {
	// The email message must be in a valid state before it can be published.
	if err := message.ValidateData(); err != nil {
		return err
	}

	// Publishing the request delegates the responsibility of delivery
	// to an asynchronous mechanism such as a queue or task scheduler.
	if err := re.Publisher.Publish(message); err != nil {
		return fmt.Errorf("failed to request email sending: %w", err)
	}

	return nil
}

