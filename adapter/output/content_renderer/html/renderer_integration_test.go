package htmlrenderer_test

import (
	"emailservice/core/application/apperrors"
	"emailservice/core/application/message"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRender_TemplateNotFoundError(t *testing.T) {
	message := message.Message{
		Id:        "123456",
		Type:      "incorrect",
		To:        "email@email.com",
		Variables: nil,
	}
	subject, html, err := rendererAdapter.Render(message)

	require.Equal(t, "", html)
	require.Equal(t, "", subject)
	require.Error(t, err)

	var targetError *apperrors.InfrastructureError
	if assert.ErrorAs(t, err, &targetError) {
		assert.Contains(t, targetError.Message, "template not found")
		assert.Nil(t, targetError.OriginalCause)
	}
}

func TestRender_RenderError(t *testing.T) {
	message := message.Message{
		Id:   "123456",
		Type: message.MessageTypeAccountDeletionCode,
		To:   "email@email.com",
		Variables: map[string]any{
			"key_not_exists": 10,
		},
	}
	subject, html, err := rendererAdapter.Render(message)

	require.Error(t, err)
	require.Equal(t, "", subject)
	require.Equal(t, "", html)

	var targetError *apperrors.InfrastructureError
	if assert.ErrorAs(t, err, &targetError) {
		assert.Contains(t, targetError.Message, "failed to render email template")
		assert.Error(t, targetError.OriginalCause)
	}
}

func TestRender_EmailVerificationCodeTemplate_Success(t *testing.T) {
	message, _ := message.NewMessage(
		"fake-id",
		"user@test.com",
		message.MessageTypeEmailVerificationCode,
		variables,
	)

	subject, html, err := rendererAdapter.Render(*message)

	require.NoError(t, err)
	assert.NotEmpty(t, html)
	assert.NotEmpty(t, subject)
	assert.Contains(t, html, "123456")
}

func TestRender_NotifyEmailVerifiedTemplate_Success(t *testing.T) {
	message, _ := message.NewMessage(
		"fake-id",
		"user@test.com",
		message.MessageTypeNotifyEmailVerified,
		nil,
	)

	subject, html, err := rendererAdapter.Render(*message)

	require.NoError(t, err)
	assert.NotEmpty(t, subject)
	assert.NotEmpty(t, html)
}

func TestRender_ChangeEmailCodeTemplate_Success(t *testing.T) {
	message, _ := message.NewMessage(
		"fake-id",
		"user@test.com",
		message.MessageTypeChangeEmailCode,
		variables,
	)

	subject, html, err := rendererAdapter.Render(*message)

	require.NoError(t, err)
	assert.NotEmpty(t, subject)
	assert.NotEmpty(t, html)
	assert.Contains(t, html, "123456")
}

func TestRender_NotifyEmailChangedTemplate_Success(t *testing.T) {
	message, _ := message.NewMessage(
		"fake-id",
		"user@test.com",
		message.MessageTypeNotifyEmailChanged,
		nil,
	)

	subject, html, err := rendererAdapter.Render(*message)

	require.NoError(t, err)
	assert.NotEmpty(t, subject)
	assert.NotEmpty(t, html)

}

func TestRender_ChangePasswordCodeTemplate_Success(t *testing.T) {
	message, _ := message.NewMessage(
		"fake-id",
		"user@test.com",
		message.MessageTypeResetPasswordCode,
		variables,
	)

	subject, html, err := rendererAdapter.Render(*message)

	require.NoError(t, err)
	assert.NotEmpty(t, subject)
	assert.NotEmpty(t, html)
	assert.Contains(t, html, "123456")
}

func TestRender_NotifyPasswordChangedTemplate_Success(t *testing.T) {
	message, _ := message.NewMessage(
		"fake-id",
		"user@test.com",
		message.MessageTypeNotifyPasswordChanged,
		nil,
	)

	subject, html, err := rendererAdapter.Render(*message)

	require.NoError(t, err)
	assert.NotEmpty(t, subject)
	assert.NotEmpty(t, html)
}

func TestRender_ResetPasswordCodeTemplate_Success(t *testing.T) {
	message, _ := message.NewMessage(
		"fake-id",
		"user@test.com",
		message.MessageTypeResetPasswordCode,
		variables,
	)

	subject, html, err := rendererAdapter.Render(*message)

	require.NoError(t, err)
	assert.NotEmpty(t, subject)
	assert.NotEmpty(t, html)
	assert.Contains(t, html, "123456")
}

func TestRender_NotifyPasswordResetTemplate_Success(t *testing.T) {
	message, _ := message.NewMessage(
		"fake-id",
		"user@test.com",
		message.MessageTypeNotifyPasswordReset,
		nil,
	)

	subject, html, err := rendererAdapter.Render(*message)

	require.NoError(t, err)
	assert.NotEmpty(t, subject)
	assert.NotEmpty(t, html)
}

func TestRender_AccountDeletionCodeTemplate_Success(t *testing.T) {
	message, _ := message.NewMessage(
		"fake-id",
		"user@test.com",
		message.MessageTypeAccountDeletionCode,
		variables,
	)

	subject, html, err := rendererAdapter.Render(*message)

	require.NoError(t, err)
	assert.NotEmpty(t, subject)
	assert.NotEmpty(t, html)
	assert.Contains(t, html, "123456")
}

func TestRender_NotifyAccountDeletedTemplate_Success(t *testing.T) {
	message, _ := message.NewMessage(
		"fake-id",
		"user@test.com",
		message.MessageTypeNotifyAccountDeleted,
		nil,
	)

	subject, html, err := rendererAdapter.Render(*message)

	require.NoError(t, err)
	assert.NotEmpty(t, subject)
	assert.NotEmpty(t, html)
}
