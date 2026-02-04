package outputport

import "emailservice/core/application/email_message"

// RenderEmailContentOutputPort defines the output port responsible for rendering
// the email content based on an EmailMessage.
//
// Implementations of this interface are expected to convert an EmailMessage
// into a rendered representation, typically HTML, suitable for sending
// through an email delivery provider.
type RenderEmailContentOutputPort interface {
	// Render converts the given EmailMessage into its final rendered content.
	// The returned string represents the rendered body, and an error is returned
	// if the content cannot be generated.
	Render(message emailmessage.EmailMessage) (string, error)
}
