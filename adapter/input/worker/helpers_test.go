package worker_test

import (
	"emailservice/adapter/input/worker"
	"emailservice/core/application/email_message"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func mustMarshal(t *testing.T, v any) []byte {
	t.Helper()
	b, err := json.Marshal(v)
	require.NoError(t, err)
	return b
}

func TestToEmailMessage_UnknownType(t *testing.T) {
	payload := mustMarshal(t, map[string]any{
		"To":        "user@test.com",
		"Subject":   "Unknown",
		"EmailType": "invalid_type",
		"BodyData":  map[string]any{},
	})

	msg, err := worker.ToEmailMessage(payload)

	require.Error(t, err)
	require.Nil(t, msg)
}

func TestToEmailMessage_EmailVerificationCode(t *testing.T) {
	var body emailmessage.EmailVerificationCodeBody
	body.VerificationCode = "123456"
	body.CodeExpirationTime = "2"
	body.EmailVerificationDeadlineDays = "3"

	payload := mustMarshal(t, map[string]any{
		"To":        "user@test.com",
		"Subject":   "Verify",
		"EmailType": emailmessage.EmailTypeEmailVerificationCode,
		"BodyData":  body,
	})

	msg, err := worker.ToEmailMessage(payload)

	require.NoError(t, err)
	require.NotNil(t, msg)
	require.Equal(t, "user@test.com", msg.GetTo())
	require.Equal(t, "Verify", msg.GetSubject())
	require.Equal(t, emailmessage.EmailTypeEmailVerificationCode, msg.GetEmailType())

	result := msg.GetBodyData().(emailmessage.EmailVerificationCodeBody)
	require.Equal(t, "123456", result.VerificationCode)
	require.Equal(t, "2", result.CodeExpirationTime)
	require.Equal(t, "3", result.EmailVerificationDeadlineDays)
}

func TestToEmailMessage_NotifyEmailVerification(t *testing.T) {
	var body emailmessage.NotifyEmailVerificationBody

	payload := mustMarshal(t, map[string]any{
		"To":        "user@test.com",
		"Subject":   "Verified",
		"EmailType": emailmessage.EmailTypeNotifyEmailVerification,
		"BodyData":  body,
	})

	msg, err := worker.ToEmailMessage(payload)

	require.NoError(t, err)
	require.NotNil(t, msg)
	require.Equal(t, "user@test.com", msg.GetTo())
	require.Equal(t, "Verified", msg.GetSubject())
	require.Equal(t, emailmessage.EmailTypeNotifyEmailVerification, msg.GetEmailType())
}

func TestToEmailMessage_ChangeEmailCode(t *testing.T) {
	var body emailmessage.ChangeEmailCodeBody
	body.VerificationCode = "999999"
	body.CodeExpirationTime = "1"

	payload := mustMarshal(t, map[string]any{
		"To":        "user@test.com",
		"Subject":   "Change Email",
		"EmailType": emailmessage.EmailTypeChangeEmailCode,
		"BodyData":  body,
	})

	msg, err := worker.ToEmailMessage(payload)

	require.NoError(t, err)
	require.NotNil(t, msg)
	require.Equal(t, "user@test.com", msg.GetTo())
	require.Equal(t, "Change Email", msg.GetSubject())
	require.Equal(t, emailmessage.EmailTypeChangeEmailCode, msg.GetEmailType())

	result := msg.GetBodyData().(emailmessage.ChangeEmailCodeBody)
	require.Equal(t, "999999", result.VerificationCode)
	require.Equal(t, "1", result.CodeExpirationTime)
}

func TestToEmailMessage_NotifyChangeEmail(t *testing.T) {
	var body emailmessage.NotifyChangeEmailBody

	payload := mustMarshal(t, map[string]any{
		"To":        "user@test.com",
		"Subject":   "Email Changed",
		"EmailType": emailmessage.EmailTypeNotifyChangeEmail,
		"BodyData":  body,
	})

	msg, err := worker.ToEmailMessage(payload)

	require.NoError(t, err)
	require.NotNil(t, msg)
	require.Equal(t, "user@test.com", msg.GetTo())
	require.Equal(t, "Email Changed", msg.GetSubject())
	require.Equal(t, emailmessage.EmailTypeNotifyChangeEmail, msg.GetEmailType())
}

func TestToEmailMessage_ResetPasswordCode(t *testing.T) {
	var body emailmessage.ResetPasswordCodeBody
	body.VerificationCode = "555555"
	body.CodeExpirationTime = "4"

	payload := mustMarshal(t, map[string]any{
		"To":        "user@test.com",
		"Subject":   "Reset Password",
		"EmailType": emailmessage.EmailTypeResetPasswordCode,
		"BodyData":  body,
	})

	msg, err := worker.ToEmailMessage(payload)

	require.NoError(t, err)
	require.NotNil(t, msg)
	require.Equal(t, "user@test.com", msg.GetTo())
	require.Equal(t, "Reset Password", msg.GetSubject())
	require.Equal(t, emailmessage.EmailTypeResetPasswordCode, msg.GetEmailType())

	result := msg.GetBodyData().(emailmessage.ResetPasswordCodeBody)
	require.Equal(t, "555555", result.VerificationCode)
	require.Equal(t, "4", result.CodeExpirationTime)
}

func TestToEmailMessage_NotifyResetPassword(t *testing.T) {
	var body emailmessage.NotifyResetPasswordBody

	payload := mustMarshal(t, map[string]any{
		"To":        "user@test.com",
		"Subject":   "Password Reset",
		"EmailType": emailmessage.EmailTypeNotifyResetPassword,
		"BodyData":  body,
	})

	msg, err := worker.ToEmailMessage(payload)

	require.NoError(t, err)
	require.NotNil(t, msg)
	require.Equal(t, "user@test.com", msg.GetTo())
	require.Equal(t, "Password Reset", msg.GetSubject())
	require.Equal(t, emailmessage.EmailTypeNotifyResetPassword, msg.GetEmailType())
}

func TestToEmailMessage_ChangePasswordCode(t *testing.T) {
	var body emailmessage.ChangePasswordCodeBody
	body.VerificationCode = "777777"
	body.CodeExpirationTime = "2"

	payload := mustMarshal(t, map[string]any{
		"To":        "user@test.com",
		"Subject":   "Change Password",
		"EmailType": emailmessage.EmailTypeChangePasswordCode,
		"BodyData":  body,
	})

	msg, err := worker.ToEmailMessage(payload)

	require.NoError(t, err)
	require.NotNil(t, msg)
	require.Equal(t, "user@test.com", msg.GetTo())
	require.Equal(t, "Change Password", msg.GetSubject())
	require.Equal(t, emailmessage.EmailTypeChangePasswordCode, msg.GetEmailType())

	result := msg.GetBodyData().(emailmessage.ChangePasswordCodeBody)
	require.Equal(t, "777777", result.VerificationCode)
	require.Equal(t, "2", result.CodeExpirationTime)
}

func TestToEmailMessage_NotifyChangePassword(t *testing.T) {
	var body emailmessage.NotifyChangePasswordBody

	payload := mustMarshal(t, map[string]any{
		"To":        "user@test.com",
		"Subject":   "Password Changed",
		"EmailType": emailmessage.EmailTypeNotifyChangePassword,
		"BodyData":  body,
	})

	msg, err := worker.ToEmailMessage(payload)

	require.NoError(t, err)
	require.NotNil(t, msg)
	require.Equal(t, "user@test.com", msg.GetTo())
	require.Equal(t, "Password Changed", msg.GetSubject())
	require.Equal(t, emailmessage.EmailTypeNotifyChangePassword, msg.GetEmailType())
}

func TestToEmailMessage_DeletionCode(t *testing.T) {
	var body emailmessage.DeletionCodeBody
	body.VerificationCode = "888888"
	body.CodeExpirationTime = "1"

	payload := mustMarshal(t, map[string]any{
		"To":        "user@test.com",
		"Subject":   "Account Deletion",
		"EmailType": emailmessage.EmailTypeDeletionCode,
		"BodyData":  body,
	})

	msg, err := worker.ToEmailMessage(payload)

	require.NoError(t, err)
	require.NotNil(t, msg)
	require.Equal(t, "user@test.com", msg.GetTo())
	require.Equal(t, "Account Deletion", msg.GetSubject())
	require.Equal(t, emailmessage.EmailTypeDeletionCode, msg.GetEmailType())

	result := msg.GetBodyData().(emailmessage.DeletionCodeBody)
	require.Equal(t, "888888", result.VerificationCode)
	require.Equal(t, "1", result.CodeExpirationTime)
}

func TestToEmailMessage_NotifyDeletion(t *testing.T) {
	var body emailmessage.NotifyDeletionBody

	payload := mustMarshal(t, map[string]any{
		"To":        "user@test.com",
		"Subject":   "Account Deleted",
		"EmailType": emailmessage.EmailTypeNotifyDeletion,
		"BodyData":  body,
	})

	msg, err := worker.ToEmailMessage(payload)

	require.NoError(t, err)
	require.NotNil(t, msg)
	require.Equal(t, "user@test.com", msg.GetTo())
	require.Equal(t, "Account Deleted", msg.GetSubject())
	require.Equal(t, emailmessage.EmailTypeNotifyDeletion, msg.GetEmailType())
}
