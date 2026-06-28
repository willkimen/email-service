package emailresend

import (
	"github.com/resend/resend-go/v3"
)

// NewResendEmailSenderAdapter creates and initializes a new instance of
// ResendEmailSenderAdapter ready for production use.
//
// It automatically extracts and maps the necessary sub-services (like client.Emails)
// from the concrete Resend Client into the adapter's internal interface.
//
// Parameters:
//   - client: The concrete *resend.Client instance provided by the official Resend SDK.
//   - from: The default sender email address configured for all outgoing communication.
//
// Returns:
//   - A pointer to the configured and ready-to-use ResendEmailSenderAdapter.
func NewResendEmailSenderAdapter(
	client *resend.Client,
	from string,
) *ResendEmailSenderAdapter {
	return &ResendEmailSenderAdapter{
		Emails: client.Emails,
		From:   from,
	}
}
