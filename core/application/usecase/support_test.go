package usecase_test

import (
	"emailservice/core/application/message"
	"errors"
)

// =================== Email message =============================

var originalCause = errors.New("Internal error")

var messageDummy, _ = message.NewMessage(
	"id", "email@email.com", message.MessageTypeNotifyAccountDeleted, nil,
)

// ================== Fake publisher ==========================
type fakePublisher struct {
	Err error
}

func (p fakePublisher) Publish(message message.Message) error {
	return p.Err
}

// =================== Fake renderer ==============================
type fakeRenderer struct {
	Subject string
	Body    string
	Err     error
}

func (f fakeRenderer) Render(message message.Message) (string, string, error) {
	return f.Subject, f.Body, f.Err
}

// ==================== Fake sender ==========================
type fakeSender struct {
	Err error
}

func (f fakeSender) SendEmail(to, subject, body string) error {
	return f.Err
}
