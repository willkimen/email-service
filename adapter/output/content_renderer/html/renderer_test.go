package renderer

import (
	"emailservice/core/application/email_message"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRender_WhenEmailTypeDoesNotExist_ShouldReturnError(t *testing.T) {
	emailMessage := FakeEmailMessageWithEmailTypeNotExist{}
	_, err := renderer.Render(emailMessage)

	assert.Error(t, err)
	assert.ErrorContains(t, err, "template not found")
}

func TestRender_WhenTemplateDataIsInvalid_ShouldReturnError(t *testing.T) {
	invalidMessage := FakeEmailMessageWithDataInvalid{FieldNotExist: "not_exist"}
	_, err := renderer.Render(invalidMessage)

	assert.Error(t, err)
	assert.ErrorContains(t, err, "failed to execute email template")
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

	html, err := renderer.Render(message)

	assert.NoError(t, err)
	assert.NotEmpty(t, html)
	assert.Contains(t, html, "123456")
	assert.Contains(t, html, "https://example.com/activate")
	assert.Contains(t, html, "7 days")
	assert.Contains(t, html, "2 hours")
}

func TestRender_ShouldRender_NotifyActivationTemplate_Correctly(t *testing.T) {
	message := emailmessage.NewNotifiyActivation(
		"user@test.com",
		"subject-test",
		"https://example.com/login",
	)

	html, err := renderer.Render(message)

	assert.NoError(t, err)
	assert.NotEmpty(t, html)
	assert.Contains(t, html, "https://example.com/login")
}

func TestRender_ShouldRender_ChangeEmailCodeTemplate_Correctly(t *testing.T) {
	message := emailmessage.NewChangeEmailCode(
		"user@test.com",
		"subject-test",
		"123456",
		"2",
	)

	html, err := renderer.Render(message)

	assert.NoError(t, err)
	assert.NotEmpty(t, html)
	assert.Contains(t, html, "123456")
	assert.Contains(t, html, "2 hours")
}

func TestRender_ShouldRender_NotifyChangeEmailTemplate_Correctly(t *testing.T) {
	message := emailmessage.NewNotifyChangeEmail(
		"user@test.com",
		"subject-test",
		"https://example.com/login",
	)

	html, err := renderer.Render(message)

	assert.NoError(t, err)
	assert.NotEmpty(t, html)
	assert.Contains(t, html, "https://example.com/login")
}

func TestRender_ShouldRender_ChangePasswordCodeTemplate_Correctly(t *testing.T) {
	message := emailmessage.NewChangePasswordCode(
		"user@test.com",
		"subject-test",
		"123456",
		"2",
	)

	html, err := renderer.Render(message)

	assert.NoError(t, err)
	assert.NotEmpty(t, html)
	assert.Contains(t, html, "123456")
	assert.Contains(t, html, "2 hours")
}

func TestRender_ShouldRender_NotifyChangePasswordTemplate_Correctly(t *testing.T) {
	message := emailmessage.NewNotifyChangePassword(
		"user@test.com",
		"subject-test",
		"https://example.com/login",
	)

	html, err := renderer.Render(message)

	assert.NoError(t, err)
	assert.NotEmpty(t, html)
	assert.Contains(t, html, "https://example.com/login")
}

func TestRender_ShouldRender_ResetPasswordCodeTemplate_Correctly(t *testing.T) {
	message := emailmessage.NewResetPasswordCode(
		"user@test.com",
		"subject-test",
		"123456",
		"https://example.com/link",
		"2",
	)

	html, err := renderer.Render(message)

	assert.NoError(t, err)
	assert.NotEmpty(t, html)
	assert.Contains(t, html, "123456")
	assert.Contains(t, html, "2 hours")
	assert.Contains(t, html, "https://example.com/link")
}

func TestRender_ShouldRender_NotifyResetPasswordTemplate_Correctly(t *testing.T) {
	message := emailmessage.NewNotifyResetPassword(
		"user@test.com",
		"subject-test",
		"https://example.com/login",
	)

	html, err := renderer.Render(message)

	assert.NoError(t, err)
	assert.NotEmpty(t, html)
	assert.Contains(t, html, "https://example.com/login")
}

func TestRender_ShouldRender_DeletionCodeTemplate_Correctly(t *testing.T) {
	message := emailmessage.NewDeletionCode(
		"user@test.com",
		"subject-test",
		"123456",
		"2",
	)

	html, err := renderer.Render(message)

	assert.NoError(t, err)
	assert.NotEmpty(t, html)
	assert.Contains(t, html, "123456")
	assert.Contains(t, html, "2 hours")
}

func TestRender_ShouldRender_NotifyDeletionTemplate_Correctly(t *testing.T) {
	message := emailmessage.NewNotifyDeletion(
		"user@test.com",
		"subject-test",
	)

	html, err := renderer.Render(message)

	assert.NoError(t, err)
	assert.NotEmpty(t, html)
}
