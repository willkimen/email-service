package htmlrenderer

import (
	"bytes"
	"emailservice/core/application/apperrors"
	"emailservice/core/application/message"
	"fmt"
	"html/template"
	"io/fs"
)

// HTMLEmailContentRendererAdapter is responsible for rendering
// the HTML body of an email.
//
// It resolves the template based on the message type exposed by
// the Message and executes it using the message variables.
type HTMLEmailContentRendererAdapter struct {
	// FS represents the file system abstraction where the HTML template
	// files reside. It allows injecting the production embedded file system
	// or custom file systems for testing purposes.
	FS fs.FS
}

// Render renders the HTML body for the given email message.
//
// The message type is used to resolve the corresponding HTML template.
// The message variables itself is passed as the template data.
//
// Returns:
//   - string: The rendered HTML body text, or an empty string ("") on error.
//   - string: The subject text, or an empty string ("") on error.
//   - error: An apperrors.InfrastructureError if any step of the process fails.
//
// Errors:
//   - Returns apperrors.InfrastructureError if no template is registered for the message type.
//   - Returns apperrors.InfrastructureError if the template cannot be parsed.
//   - Returns apperrors.InfrastructureError if failed to render email template.
func (r *HTMLEmailContentRendererAdapter) Render(message message.Message) (string, string, error) {
	path, ok := pathTemplates[message.Type]
	if !ok {
		return "", "", apperrors.NewInfrastructureError("template not found", nil)
	}

	tmpl, err := template.ParseFS(r.FS, path)
	if err != nil {
		msgError := fmt.Sprintf(
			"failed to parse email template %q",
			path,
		)
		return "", "", apperrors.NewInfrastructureError(msgError, err)
	}
	tmpl = tmpl.Option("missingkey=error")

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, message.Variables); err != nil {
		msgError := fmt.Sprintf(
			"failed to render email template %q",
			path,
		)
		return "", "", apperrors.NewInfrastructureError(msgError, err)
	}

	return subjects[message.Type], buf.String(), nil
}
