package renderer_test

import (
	"emailservice/core/application/email_message"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRender_WhenEmailTypeDoesNotExist_ShouldReturnError(t *testing.T) {
	emailMessage := FakeEmailMessageWithEmailTypeNotExist{}
	_, err := rendererAdapter.Render(emailMessage)

	require.Error(t, err,
		"expected Render to return error when email type does not exist")
	assert.ErrorContains(t, err,
		"template not found", "expected error to mention missing template")
}

func TestRender_WhenTemplateDataIsInvalid_ShouldReturnError(t *testing.T) {
	invalidMessage := FakeEmailMessageWithDataInvalid{FieldNotExist: "not_exist"}
	_, err := rendererAdapter.Render(invalidMessage)

	require.Error(t, err,
		"expected Render to return error when template data is invalid")
	assert.ErrorContains(t, err,
		"failed to execute email template", "expected error to mention template execution failure")
}

func TestRender_ShouldRender_ActivationCodeTemplate_Correctly(t *testing.T) {
	message := emailmessage.NewActivationCode(
		"user@test.com",
		"subject-test",
		"123456",
		"https://example.com/activate",
		"2",
		"7",
	)

	html, err := rendererAdapter.Render(message)

	require.NoError(t, err,
		"expected Render to succeed for ActivationCode template")
	assert.NotEmpty(t, html,
		"expected rendered HTML to be non-empty")
	assert.Contains(t, html,
		"123456", "expected HTML to contain activation code")
	assert.Contains(t, html,
		"https://example.com/activate", "expected HTML to contain activation link")
	assert.Contains(t, html, "7 days",
		"expected HTML to contain expiration days information")
	assert.Contains(t, html, "2 hours",
		"expected HTML to contain expiration hours information")
}

func TestRender_ShouldRender_NotifyActivationTemplate_Correctly(t *testing.T) {
	message := emailmessage.NewNotifiyActivation(
		"user@test.com",
		"subject-test",
		"https://example.com/login",
	)

	html, err := rendererAdapter.Render(message)

	require.NoError(t, err,
		"expected Render to succeed for NotifyActivation template")
	assert.NotEmpty(t, html,
		"expected rendered HTML to be non-empty")
	assert.Contains(t, html,
		"https://example.com/login", "expected HTML to contain login link")
}

func TestRender_ShouldRender_ChangeEmailCodeTemplate_Correctly(t *testing.T) {
	message := emailmessage.NewChangeEmailCode(
		"user@test.com",
		"subject-test",
		"123456",
		"2",
	)

	html, err := rendererAdapter.Render(message)

	require.NoError(t, err,
		"expected Render to succeed for ChangeEmailCode template")
	assert.NotEmpty(t, html,
		"expected rendered HTML to be non-empty")
	assert.Contains(t, html,
		"123456", "expected HTML to contain change email code")
	assert.Contains(t, html, "2 hours",
		"expected HTML to contain expiration hours information")
}

func TestRender_ShouldRender_NotifyChangeEmailTemplate_Correctly(t *testing.T) {
	message := emailmessage.NewNotifyChangeEmail(
		"user@test.com",
		"subject-test",
		"https://example.com/login",
	)

	html, err := rendererAdapter.Render(message)

	require.NoError(t, err,
		"expected Render to succeed for NotifyChangeEmail template")
	assert.NotEmpty(t, html,
		"expected rendered HTML to be non-empty")
	assert.Contains(t, html, "https://example.com/login",
		"expected HTML to contain login link")
}

func TestRender_ShouldRender_ChangePasswordCodeTemplate_Correctly(t *testing.T) {
	message := emailmessage.NewChangePasswordCode(
		"user@test.com",
		"subject-test",
		"123456",
		"2",
	)

	html, err := rendererAdapter.Render(message)

	require.NoError(t, err,
		"expected Render to succeed for ChangePasswordCode template")
	assert.NotEmpty(t, html,
		"expected rendered HTML to be non-empty")
	assert.Contains(t, html, "123456",
		"expected HTML to contain change password code")
	assert.Contains(t, html, "2 hours",
		"expected HTML to contain expiration hours information")
}

func TestRender_ShouldRender_NotifyChangePasswordTemplate_Correctly(t *testing.T) {
	message := emailmessage.NewNotifyChangePassword(
		"user@test.com",
		"subject-test",
		"https://example.com/login",
	)

	html, err := rendererAdapter.Render(message)

	require.NoError(t, err,
		"expected Render to succeed for NotifyChangePassword template")
	assert.NotEmpty(t, html,
		"expected rendered HTML to be non-empty")
	assert.Contains(t, html,
		"https://example.com/login", "expected HTML to contain login link")
}

func TestRender_ShouldRender_ResetPasswordCodeTemplate_Correctly(t *testing.T) {
	message := emailmessage.NewResetPasswordCode(
		"user@test.com",
		"subject-test",
		"123456",
		"https://example.com/link",
		"2",
	)

	html, err := rendererAdapter.Render(message)

	require.NoError(t, err,
		"expected Render to succeed for ResetPasswordCode template")
	assert.NotEmpty(t, html,
		"expected rendered HTML to be non-empty")
	assert.Contains(t, html,
		"123456", "expected HTML to contain reset password code")
	assert.Contains(t, html,
		"2 hours", "expected HTML to contain expiration hours information")
	assert.Contains(t, html,
		"https://example.com/link", "expected HTML to contain reset password link")
}

func TestRender_ShouldRender_NotifyResetPasswordTemplate_Correctly(t *testing.T) {
	message := emailmessage.NewNotifyResetPassword(
		"user@test.com",
		"subject-test",
		"https://example.com/login",
	)

	html, err := rendererAdapter.Render(message)

	require.NoError(t, err,
		"expected Render to succeed for NotifyResetPassword template")
	assert.NotEmpty(t, html,
		"expected rendered HTML to be non-empty")
	assert.Contains(t, html, "https://example.com/login",
		"expected HTML to contain login link")
}

func TestRender_ShouldRender_DeletionCodeTemplate_Correctly(t *testing.T) {
	message := emailmessage.NewDeletionCode(
		"user@test.com",
		"subject-test",
		"123456",
		"2",
	)

	html, err := rendererAdapter.Render(message)

	require.NoError(t, err,
		"expected Render to succeed for DeletionCode template")
	assert.NotEmpty(t, html,
		"expected rendered HTML to be non-empty")
	assert.Contains(t, html, "123456",
		"expected HTML to contain deletion code")
	assert.Contains(t, html, "2 hours",
		"expected HTML to contain expiration hours information")
}

func TestRender_ShouldRender_NotifyDeletionTemplate_Correctly(t *testing.T) {
	message := emailmessage.NewNotifyDeletion(
		"user@test.com",
		"subject-test",
	)

	html, err := rendererAdapter.Render(message)

	require.NoError(t, err,
		"expected Render to succeed for NotifyDeletion template")
	assert.NotEmpty(t, html,
		"expected rendered HTML to be non-empty")
}
