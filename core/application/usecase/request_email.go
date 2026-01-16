package usecase

import (
	"emailservice/core/application/email_message"
	"emailservice/core/application/ports/output"
	"fmt"
)

type RequestEmail struct {
	Publisher outputport.EmailRequestPublisher
}

func (re *RequestEmail) Request(message emailmessage.EmailMessage) error {
	err := message.ValidateData()

	if err != nil {
		return err

	}

	err = re.Publisher.Publish(message)

	if err != nil {
		return fmt.Errorf("failed to request email sending: %w", err)
	}

	return nil
}
