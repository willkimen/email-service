package renderer

import (
	"bytes"
	"emailservice/core/application/email_message"
	"errors"
	"fmt"
	"html/template"
)

// HTMLEmailContentRendererAdapter is responsible for rendering
// the HTML body of an email message.
//
// It resolves the template based on the email type exposed by
// the EmailMessage and executes it using the message data.
type HTMLEmailContentRendererAdapter struct {
}

func NewHTMLEmailContentRendererAdapter() *HTMLEmailContentRendererAdapter {
	return &HTMLEmailContentRendererAdapter{}
}

// Render renders the HTML body for the given email message.
//
// The email type returned by GetEmailType is used to resolve
// the corresponding HTML template. The message itself is passed
// as the template data.
//
// An error is returned if no template is registered for the email
// type, if the template cannot be parsed, or if execution fails.
func (r *HTMLEmailContentRendererAdapter) Render(message emailmessage.EmailMessage) (string, error) {
	path, ok := pathTemplates[message.GetEmailType()]
	if !ok {
		return "", errors.New("template not found")
	}

	tmpl, err := template.ParseFiles(path)
	if err != nil {
		return "", fmt.Errorf(
			"failed to parse email template %q: %w",
			path,
			err,
		)
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, message); err != nil {
		return "", fmt.Errorf(
			"failed to execute email template %q: %w",
			path,
			err,
		)
	}

	return buf.String(), nil
}
