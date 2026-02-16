//go:build email

package emailsender_test

import (
	"emailservice/adapter/output/resend"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/require"
)

func TestResendEmailSenderAdapter_SendEmail_Integration(t *testing.T) {
	// This is an integration test that verifies the adapter can successfully
	// communicate with the external email service and submit a send request.
	//
	// It does NOT validate email delivery outcomes (delivered, bounced, spam, etc.),
	// since those results are handled asynchronously via webhooks.
	// The goal here is only to ensure that configuration, authentication,
	// and request formatting are correct.
	err := godotenv.Load("../../../.env")
	require.NoError(
		t,
		err,
		"expected .env file to load successfully for integration test",
	)

	require.NotEmpty(
		t,
		os.Getenv("RESEND_API_KEY"),
		"expected RESEND_API_KEY to be set in environment variables",
	)

	require.NotEmpty(
		t,
		os.Getenv("FROM_EMAIL"),
		"expected FROM_EMAIL to be set in environment variables",
	)

	adapter := &emailsender.ResendEmailSenderAdapter{}

	err = adapter.SendEmail(
		"delivered@resend.dev",
		"Integration test email",
		"<p>Integration test</p>",
	)

	require.NoError(
		t,
		err,
		"expected SendEmail to complete without returning an error",
	)
}
