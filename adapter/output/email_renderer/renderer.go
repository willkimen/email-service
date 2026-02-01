package renderer

import (
	"bytes"
	"emailservice/core/application/email_message"
	"errors"
	"fmt"
	"html/template"
)

// HTMLEmailRenderer is responsible for rendering email bodies from HTML templates.
//
// It resolves a template path based on a template identifier and executes
// the template using the provided data, returning the rendered HTML as a string.
type HTMLEmailRenderer struct {
}

// Render renders an HTML email template identified by templateID using the given data.
//
// It looks up the template path, parses the HTML file, and executes it.
// An error is returned if the template does not exist or if rendering fails.
func (r *HTMLEmailRenderer) Render(message emailmessage.EmailMessage) (string, error) {
	path, ok := pathTemplates[message.TemplateID()]
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
