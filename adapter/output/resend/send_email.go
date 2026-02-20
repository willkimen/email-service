// Package emailsender provides adapters responsible for sending emails
// through external email delivery services.
//
// This package belongs to the infrastructure layer and implements
// outbound communication with third-party providers, translating
// application-level intent into concrete delivery operations.
package emailsender

import (
	"emailservice/core/application/email_errors"
	"errors"
	"fmt"

	"github.com/resend/resend-go/v3"
)

// resendEmailAPI abstracts the subset of the Resend client
// required by this adapter.
//
// This interface exists to decouple the adapter from the concrete
// *resend.Client implementation and to allow mocking in tests.
type resendEmailAPI interface {
	Send(params *resend.SendEmailRequest) (*resend.SendEmailResponse, error)
}

// ResendEmailSenderAdapter is an email sender adapter that delivers
// messages using the Resend email service.
//
// It belongs to the infrastructure layer and translates application
// email requests into calls to the Resend API.
type ResendEmailSenderAdapter struct {
	Emails resendEmailAPI
	From   string
}

// SendEmail sends an email message to the given recipient with the provided
// subject and HTML body.
//
// If the email service temporarily rejects the request (for example, due to
// rate limiting), the returned error preserves the underlying cause so it can
// be classified by upper layers as retryable.
//
// Permanent failures are returned as errors without retry guarantees.
func (r *ResendEmailSenderAdapter) SendEmail(to, subject, body string) error {
	params := &resend.SendEmailRequest{
		To:      []string{to},
		From:    r.From,
		Subject: subject,
		Html:    body,
	}

	_, err := r.Emails.Send(params)
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
