package emailsender

import (
	"emailservice/core/application/email_errors"
	"errors"
	"testing"

	"github.com/resend/resend-go/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// mockResendEmailAPI is a test double used to simulate
// different behaviors of the Resend client.
type mockResendEmailAPI struct {
	sendFunc func(params *resend.SendEmailRequest) (*resend.SendEmailResponse, error)
}

func (m *mockResendEmailAPI) Send(params *resend.SendEmailRequest) (*resend.SendEmailResponse, error) {
	return m.sendFunc(params)
}

func TestSendEmail_TemporaryFailure(t *testing.T) {
	mock := &mockResendEmailAPI{
		sendFunc: func(params *resend.SendEmailRequest) (*resend.SendEmailResponse, error) {
			return nil, resend.ErrRateLimit
		},
	}

	adapter := &ResendEmailSenderAdapter{
		Emails: mock,
		From:   "test@example.com",
	}

	err := adapter.SendEmail("user@test.com", "subject", "<p>body</p>")

	require.Error(t, err, "expected an error when rate limit occurs")
	assert.ErrorIs(t, err, emailerrors.ErrTemporaryFailure,
		"expected error to wrap ErrTemporaryFailure")
}

func TestSendEmail_PermanentFailure(t *testing.T) {
	mock := &mockResendEmailAPI{
		sendFunc: func(params *resend.SendEmailRequest) (*resend.SendEmailResponse, error) {
			return nil, errors.New("some permanent failure")
		},
	}

	adapter := &ResendEmailSenderAdapter{
		Emails: mock,
		From:   "test@example.com",
	}

	err := adapter.SendEmail("user@test.com", "subject", "<p>body</p>")

	require.Error(t, err, "expected an error for permanent failure")
	assert.ErrorIs(t, err, emailerrors.ErrPermanentFailure,
		"expected error to wrap ErrPermanentFailure")
}

func TestSendEmail_Success(t *testing.T) {
	mock := &mockResendEmailAPI{
		sendFunc: func(params *resend.SendEmailRequest) (*resend.SendEmailResponse, error) {
			return &resend.SendEmailResponse{}, nil
		},
	}

	adapter := &ResendEmailSenderAdapter{
		Emails: mock,
		From:   "test@example.com",
	}

	err := adapter.SendEmail("user@test.com", "subject", "<p>body</p>")

	require.NoError(t, err, "expected no error when email is sent successfully")
}

