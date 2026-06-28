package emailresend_test

import (
	"emailservice/adapter/output/resend"
	"emailservice/core/application/apperrors"
	"testing"

	"github.com/resend/resend-go/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSendEmail_Success(t *testing.T) {
	adapter := &emailresend.ResendEmailSenderAdapter{
		Emails: mockEmailsSuccess,
		From:   "test@example.com",
	}

	err := adapter.SendEmail("user@test.com", "subject", "<p>body</p>")

	require.NoError(t, err)
}

func TestSendEmail_TemporaryFailure(t *testing.T) {
	adapter := &emailresend.ResendEmailSenderAdapter{
		Emails: mockEmailsWithTemporaryFailure,
		From:   "test@example.com",
	}

	actualErr := adapter.SendEmail("user@test.com", "subject", "<p>body</p>")

	require.Error(t, actualErr)

	var targetError *apperrors.InfrastructureError
	if assert.ErrorAs(t, actualErr, &targetError) {
		assert.ErrorIs(t, targetError.OriginalCause, apperrors.ErrTemporaryFailure)
		assert.ErrorIs(t, targetError.OriginalCause, resend.ErrRateLimit)
	}
}

func TestSendEmail_PermanentFailure(t *testing.T) {
	adapter := &emailresend.ResendEmailSenderAdapter{
		Emails: mockEmailsWithPermanentFailure,
		From:   "test@example.com",
	}

	actualErr := adapter.SendEmail("user@test.com", "subject", "<p>body</p>")

	require.Error(t, actualErr)
	var targetError *apperrors.InfrastructureError
	if assert.ErrorAs(t, actualErr, &targetError) {
		assert.ErrorIs(t, targetError.OriginalCause, apperrors.ErrPermanentFailure)
		assert.ErrorIs(t, targetError.OriginalCause, errCausePermanentFailure)
	}
}
