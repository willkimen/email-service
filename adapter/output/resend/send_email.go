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
	"log/slog"

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
	Logger *slog.Logger
}

func NewResendEmailSenderAdapter(
	client *resend.Client,
	from string,
	logger *slog.Logger,
) *ResendEmailSenderAdapter {
	return &ResendEmailSenderAdapter{
		Emails: client.Emails,
		From:   from,
		Logger: logger,
	}
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
	r.Logger.Info(
		"sending email via resend",
		"to", to,
		"subject", subject,
		"from", r.From,
	)

	params := &resend.SendEmailRequest{
		To:      []string{to},
		From:    r.From,
		Subject: subject,
		Html:    body,
	}

	resp, err := r.Emails.Send(params)
	if err != nil {
		// Rate limit errors are propagated so the application layer
		// can decide whether the operation should be retried.
		if errors.Is(err, resend.ErrRateLimit) {
			r.Logger.Error(
				"resend rate limit error",
				"error", err,
				"to", to,
				"subject", subject,
			)
			return fmt.Errorf("%w: %w", emailerrors.ErrTemporaryFailure, err)
		}

		r.Logger.Error(
			"resend permanent failure",
			"error", err,
			"to", to,
			"subject", subject,
		)
		// Other failures are treated as non-retryable by default.
		return fmt.Errorf("%w: %w", emailerrors.ErrPermanentFailure, err)
	}

	r.Logger.Info(
		"email sent successfully via resend",
		"provider_id", resp.Id,
		"to", to,
		"subject", subject,
	)

	return nil
}
