package message

// NewMessage creates and validates a new Message.
//
// It is strongly recommended to use this factory function (or workflow-specific
// factories) instead of instantiating the Message struct directly, as it centralizes
// and encapsulates all business validation rules required for a Message entity.
//
// Any provided message type must strictly match one of the system's pre-existing
// constants (e.g., MessageTypeEmailVerificationCode).
//
// Example:
//
//	variables := map[string]any{"code": "123456"}
//	msg, err := message.NewMessage(
//		"msg-123",
//		"user@example.com",
//		message.MessageTypeEmailVerificationCode,
//		variables
//	)
//	if err != nil {
//	    // handle validation error
//	}
//
// Errors:
//   - apperrors.InvalidFieldError: if any required argument is empty,
//     the email format is invalid, or the type does not match any pre-existing system type.
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
