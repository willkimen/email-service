package outputport

import "emailservice/core/application/message"

// RenderEmailContentPort defines the output port responsible for rendering
// the email content based on a Message.
//
// Implementations of this interface are expected to convert a Message
// into a rendered representation, typically HTML, suitable for sending
// through an email delivery provider.
type RenderEmailContentPort interface {
	// Render converts the given Message into its final rendered content.
	// The returned string represents the rendered body, and an error is returned
	// if the content cannot be generated.
	// Returns:
	//   - string: The rendered HTML body text, or an empty string ("") on error.
	//   - string: The subject text, or an empty string ("") on error.
	//   - error: An apperrors.InfrastructureError if any step of the process fails.
	//
	// Errors:
	//  - Returns apperrors.InfrastructureError if no template is registered for the message type.
	//  - Returns apperrors.InfrastructureError if the template cannot be parsed.
	//  - Returns apperrors.InfrastructureError if failed to render email template.
	Render(message message.Message) (string, string, error)
}
