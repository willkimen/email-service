package message

// Message represents an email message consumed by the email service.
//
// It contains the information required to identify the message, determine
// which email workflow to execute, identify the recipient, and populate
// the corresponding email template.
//
// Message is a transport-level representation of an email request received
// from the messaging broker.
type Message struct {
	// Id is the unique identifier of the message. Useful for tracing,
	// auditing, and ensuring idempotency during queue processing.
	Id string

	// Type defines the specific workflow or email template (e.g., "email_verification_code",
	// "password_reset"). Used by the consumer to determine how to process the Variables.
	Type string

	// To is the email address of the primary recipient (e.g., "user@email.com").
	To string

	// Variables contains the dynamic and variable data required to populate
	// the email template (e.g., {"name": "John", "code": "123456"}).
	Variables map[string]any
}
