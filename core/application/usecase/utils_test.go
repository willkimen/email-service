package usecase_test

import (
	"emailservice/core/application/email_message"
	"errors"
)

type publishFailureFake struct{}

func (publishFailureFake) Publish(message emailmessage.EmailMessage) error {
	return errors.New("fake error")
}

type publisherSuccessFake struct{}

func (publisherSuccessFake) Publish(message emailmessage.EmailMessage) error {
	return nil

}

var invalidMessage *emailmessage.ActivationCode = emailmessage.NewActivationCode(
	"fake@fake.com", "", "fake", "fake", "fake", "fake",
)

var messageCorrect *emailmessage.ActivationCode = emailmessage.NewActivationCode(
	"fake@fake.com", "fake", "fake", "fake", "fake", "fake",
)
