//go:build emailtest

package emailsender

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
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
	require.NoError(t, err)

	require.NotEmpty(t, os.Getenv("RESEND_API_KEY"))
	require.NotEmpty(t, os.Getenv("FROM_EMAIL"))

	adapter := &ResendEmailSenderAdapter{}

	emailID, err := adapter.SendEmail(
		"delivered@resend.dev",
		"Integration test email",
		"<p>Integration test</p>",
	)

	require.NoError(t, err)
	assert.NotEmpty(t, emailID)
}
