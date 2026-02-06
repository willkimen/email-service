// Package emailmessage defines the core representations
// used to describe email messages and their validation rules.
package emailmessage

// EmailMessage defines the behavior required by any email message
// that can be processed by the application.
//
// Implementations must validate whether the message data
// represents an acceptable and consistent state.
type EmailMessage interface {
	ValidateData() error
	TemplateID() string
	GetTo() string
	GetSubject() string
}

// BaseMessage represents the common data required by all email messages.
//
// It defines the recipient and subject used when sending emails.
type BaseMessage struct {
	To      string
	Subject string
}

func (b *BaseMessage) GetTo() string {
	return b.To
}

func (b *BaseMessage) GetSubject() string {
	return b.Subject
}

// BaseCodeMessage represents shared data for email messages that rely
// on verification codes.
//
// It defines the code itself and how long the code remains valid.
type BaseCodeMessage struct {
	VerificationCode    string
	CodeExpirationHours string
}
