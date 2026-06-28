/*
Package emailresend provides an infrastructure adapter for sending emails
using the Resend API service.

It integrates with the official Resend Go SDK, translating the application's domain-specific
email requests into concrete API payloads, managing rate limits, and handling
network resilient errors.

RECOMMENDATION: Always use the factory function NewResendEmailSenderAdapter to instantiate
the adapter instead of initializing the struct directly.

# Usage Example

	client := resend.NewClient("re_your_api_key")
	fromAddress := "noreply@yourdomain.com"

	// Recommended approach for production:
	sender := emailresend.NewResendEmailSenderAdapter(client, fromAddress)

	err := sender.SendEmail("user@example.com", "Welcome!", "<h1>Hello World</h1>")
*/
package emailresend
