// Package emailmessage defines the core representations
// used to describe email messages and their validation rules.
package emailmessage

// EmailMessage defines the behavior required by any email message
// that can be processed by the application.
//
// It represents a domain-level email abstraction, decoupled from
// rendering, transport, or template concerns.
type EmailMessage interface {
	// ValidateData validates whether the message data represents
	// an acceptable and consistent state for sending.
	ValidateData() error

	// GetEmailType returns the domain-level email type associated
	// with this message.
	//
	// The returned value is used by other layers (such as renderers
	// and adapters) to resolve templates and processing logic without
	// relying on hardcoded strings.
	GetEmailType() string

	// GetTo returns the recipient email address.
	GetTo() string

	// GetSubject returns the email subject.
	GetSubject() string

	// GetBodyData returns the data structure used to render
	// the email body template.
	GetBodyData() any
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
