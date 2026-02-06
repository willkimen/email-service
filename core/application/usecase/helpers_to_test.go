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

type FakeRenderer struct {
	Body string
	Err  error
}

func (f FakeRenderer) Render(message emailmessage.EmailMessage) (string, error) {
	return f.Body, f.Err
}

type FakeSender struct {
	Err error
}

func (f FakeSender) SendEmail(to, subject, body string) error {
	return f.Err
}

type FakeEmailMessage struct{}

func (FakeEmailMessage) ValidateData() error { return nil }
func (FakeEmailMessage) TemplateID() string  { return "anytemplate" }
func (FakeEmailMessage) GetTo() string       { return "to" }
func (FakeEmailMessage) GetSubject() string  { return "subject" }
