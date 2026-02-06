// Package emailsender provides adapters responsible for sending emails
// through external email delivery services.
//
// This package belongs to the infrastructure layer and implements
// outbound communication with third-party providers, translating
// application-level intent into concrete delivery operations.
package emailsender

import (
	emailerrors "emailservice/core/application/email_errors"
	"errors"
	"fmt"
	"os"

	"github.com/resend/resend-go/v3"
)

// ResendEmailSenderAdapter is an email sender adapter that delivers
// messages using the Resend email service.
//
// It translates application email requests into calls to the Resend API.
type ResendEmailSenderAdapter struct{}

// SendEmail sends an email message to the given recipient with the provided
// subject and HTML body.
//
// If the email service temporarily rejects the request (for example, due to
// rate limiting), the returned error preserves the underlying cause so it can
// be classified by upper layers as retryable.
//
// Permanent failures are returned as errors without retry guarantees.
func (r *ResendEmailSenderAdapter) SendEmail(to, subject, body string) error {
	client := resend.NewClient(os.Getenv("RESEND_API_KEY"))

	params := &resend.SendEmailRequest{
		To:      []string{to},
		From:    os.Getenv("FROM_EMAIL"),
		Subject: subject,
		Html:    body,
	}

	_, err := client.Emails.Send(params)
	if err != nil {
		// Rate limit errors are propagated so the application layer
		// can decide whether the operation should be retried.
		if errors.Is(err, resend.ErrRateLimit) {
			return fmt.Errorf("%w: %w", emailerrors.ErrTemporaryFailure, err)

		}

		// Other failures are treated as non-retryable by default.
		return fmt.Errorf("%w: %w", emailerrors.ErrPermanentFailure, err)

	}

	return nil
}
