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

var invalidMessage *emailmessage.EmailVerificationCode = emailmessage.NewEmailVerificationCode(
	"fake@fake.com", "", "fake", "fake", "fake", "fake",
)

var messageCorrect *emailmessage.EmailVerificationCode = emailmessage.NewEmailVerificationCode(
	"fake@fake.com", "fake", "fake", "fake", "fake", "fake",
)

type fakeRenderer struct {
	Body string
	Err  error
}

func (f fakeRenderer) Render(message emailmessage.EmailMessage) (string, error) {
	return f.Body, f.Err
}

type fakeSender struct {
	Err error
}

func (f fakeSender) SendEmail(to, subject, body string) error {
	return f.Err
}

type fakeLogger struct {
}

func (fakeLogger) Info(msg string, fields ...any)             {}
func (fakeLogger) Error(msg string, err error, fields ...any) {}

type fakeEmailMessage struct{}

func (fakeEmailMessage) ValidateData() error  { return nil }
func (fakeEmailMessage) GetEmailType() string { return "anytype" }
func (fakeEmailMessage) GetTo() string        { return "to" }
func (fakeEmailMessage) GetSubject() string   { return "subject" }
func (fakeEmailMessage) GetBodyData() any     { return nil }
