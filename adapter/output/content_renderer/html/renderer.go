package renderer

import (
	"bytes"
	"emailservice/core/application/email_message"
	"embed"
	"errors"
	"fmt"
	"html/template"
	"log/slog"
)

//go:embed templates/*.html
var templatesFS embed.FS

// HTMLEmailContentRendererAdapter is responsible for rendering
// the HTML body of an email message.
//
// It resolves the template based on the email type exposed by
// the EmailMessage and executes it using the message data.
type HTMLEmailContentRendererAdapter struct {
	Logger *slog.Logger
}

func NewHTMLEmailContentRendererAdapter(
	logger *slog.Logger,
) *HTMLEmailContentRendererAdapter {
	return &HTMLEmailContentRendererAdapter{
		Logger: logger,
	}
}

// Render renders the HTML body for the given email message.
//
// The email type returned by GetEmailType is used to resolve
// the corresponding HTML template. The message itself is passed
// as the template data.
//
// An error is returned if no template is registered for the email
// type, if the template cannot be parsed, or if execution fails.
func (r *HTMLEmailContentRendererAdapter) Render(
	message emailmessage.EmailMessage,
) (string, error) {
	emailType := message.GetEmailType()
	path, ok := pathTemplates[emailType]
	if !ok {
		r.Logger.Error(
			"email template not found",
			"email_type", emailType,
		)
		return "", errors.New("template not found")
	}

	r.Logger.Info(
		"rendering email template",
		"email_type", emailType,
		"template_path", path,
	)

	tmpl, err := template.ParseFS(templatesFS, path)
	if err != nil {
		r.Logger.Error(
			"failed to parse email template",
			"error", err,
			"email_type", emailType,
			"template_path", path,
		)

		return "", fmt.Errorf(
			"failed to parse email template %q: %w",
			path,
			err,
		)
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, message); err != nil {
		r.Logger.Error(
			"failed to execute email template",
			"error", err,
			"email_type", emailType,
			"template_path", path,
		)

		return "", fmt.Errorf(
			"failed to execute email template %q: %w",
			path,
			err,
		)
	}

	r.Logger.Info(
		"email template rendered successfully",
		"email_type", emailType,
		"template_path", path,
	)

	return buf.String(), nil
}
