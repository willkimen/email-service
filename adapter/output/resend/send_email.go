package emailresend

import (
	"emailservice/core/application/apperrors"
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
// It translates application email requests into calls to the Resend API.
type ResendEmailSenderAdapter struct {
	// Emails represents the internal API client interface used to interact
	// with Resend's services. This is abstracted via an interface to allow
	// easy mocking during unit tests.
	Emails resendEmailAPI

	// From is the default sender email address (e.g., "noreply@yourdomain.com")
	// configured for all outgoing emails processed by this adapter.
	From string
}

// SendEmail sends an email message to the given recipient with the provided
// subject and HTML body.
//
// Arguments:
//   - to: The recipient's email address.
//   - subject: The subject line of the email.
//   - body: The HTML content of the email body.
//
// Errors:
//   - An apperrors.InfrastructureError. Returns wrapped with apperrors.ErrTemporaryFailure
//     on rate limit breaches, or wrapped with apperrors.ErrPermanentFailure
//     on API or network failures.
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
			return apperrors.NewInfrastructureError(
				"send email",
				fmt.Errorf("%w: %w", apperrors.ErrTemporaryFailure, err),
			)
		}

		// Other failures are treated as non-retryable by default.
		return apperrors.NewInfrastructureError(
			"send email",
			fmt.Errorf("%w: %w", apperrors.ErrPermanentFailure, err),
		)
	}

	return nil
}
