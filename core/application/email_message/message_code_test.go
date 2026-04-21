package emailmessage_test

import (
	"emailservice/core/application/email_message"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// ============ EmailVerificationCode tests ============
func TestEmailVerificationCode_IsCreatedCorrectly(t *testing.T) {
	actualEmailVerification := validEmailVerificationCode()

	_, ok := actualEmailVerification.GetBodyData().(emailmessage.EmailVerificationCodeBody)
	require.True(t, ok, "expected body data to be of type EmailVerificationCodeBody")

	assert.Equal(t, to, actualEmailVerification.To,
		"expected To to match the provided value")
	assert.Equal(t, subject, actualEmailVerification.Subject,
		"expected Subject to match the provided value")
	assert.Equal(t, verificationCode, actualEmailVerification.VerificationCode,
		"expected VerificationCode to match the provided value")
	assert.Equal(t, link, actualEmailVerification.EmailVerificationLink,
		"expected EmailVerificationLink to match the provided value")
	assert.Equal(t, codeExpirationTime, actualEmailVerification.CodeExpirationTime,
		"expected CodeExpirationTime to match the provided value")
	assert.Equal(t, emailVerificationDeadlineDays, actualEmailVerification.EmailVerificationDeadlineDays,
		"expected EmailVerificationDeadlineDays to match the provided value")
	assert.Equal(t, emailmessage.EmailTypeEmailVerificationCode, actualEmailVerification.GetEmailType(),
		"expected email type to be EmailVerificationCode")
	assert.Nil(t, actualEmailVerification.ValidateData(),
		"expected ValidateData to return nil for a valid EmailVerificationCode")

}

func TestEmailVerificationCode_EmptyField_ReturnError(t *testing.T) {
	tests := []struct {
		name  string
		setup func(p *emailmessage.EmailVerificationCode)
	}{
		{
			name: "empty To",
			setup: func(p *emailmessage.EmailVerificationCode) {
				p.To = ""
			},
		},
		{
			name: "empty Subject",
			setup: func(p *emailmessage.EmailVerificationCode) {
				p.Subject = ""
			},
		},
		{
			name: "empty VerificationCode",
			setup: func(p *emailmessage.EmailVerificationCode) {
				p.VerificationCode = ""
			},
		},
		{
			name: "empty CodeExpirationTime",
			setup: func(p *emailmessage.EmailVerificationCode) {
				p.CodeExpirationTime = ""
			},
		},
		{
			name: "empty EmailVerificationLink",
			setup: func(p *emailmessage.EmailVerificationCode) {
				p.EmailVerificationLink = ""
			},
		},
		{
			name: "empty EmailVerificationDeadlineDays",
			setup: func(p *emailmessage.EmailVerificationCode) {
				p.EmailVerificationDeadlineDays = ""
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualEmailVerification := validEmailVerificationCode()
			tt.setup(actualEmailVerification)

			require.Error(t, actualEmailVerification.ValidateData(),
				"expected ValidateData to return an error when %s is empty", tt.name)
		})
	}
}

// ============ ChangeEmailCode tests ============
func TestChangeEmailCode_IsCreatedCorrectly(t *testing.T) {
	actualChange := validChangeEmailCode()

	_, ok := actualChange.GetBodyData().(emailmessage.ChangeEmailCodeBody)
	require.True(t, ok, "expected body data to be of type ChangeEmailCodeBody")

	assert.Equal(t, to, actualChange.To,
		"expected To to match the provided value")
	assert.Equal(t, subject, actualChange.Subject,
		"expected Subject to match the provided value")
	assert.Equal(t, verificationCode, actualChange.VerificationCode,
		"expected VerificationCode to match the provided value")
	assert.Equal(t, codeExpirationTime, actualChange.CodeExpirationTime,
		"expected CodeExpirationTime to match the provided value")
	assert.Equal(t, emailmessage.EmailTypeChangeEmailCode, actualChange.GetEmailType(),
		"expected email type to be ChangeEmailCode")
	assert.Nil(t, actualChange.ValidateData(),
		"expected ValidateData to return nil for a valid ChangeEmailCode")
}

func TestChangeEmailCode_EmptyField_ReturnError(t *testing.T) {
	tests := []struct {
		name  string
		setup func(p *emailmessage.ChangeEmailCode)
	}{
		{
			name: "empty To",
			setup: func(p *emailmessage.ChangeEmailCode) {
				p.To = ""
			},
		},
		{
			name: "empty Subject",
			setup: func(p *emailmessage.ChangeEmailCode) {
				p.Subject = ""
			},
		},
		{
			name: "empty VerificationCode",
			setup: func(p *emailmessage.ChangeEmailCode) {
				p.VerificationCode = ""
			},
		},
		{
			name: "empty CodeExpirationTime",
			setup: func(p *emailmessage.ChangeEmailCode) {
				p.CodeExpirationTime = ""
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualChange := validChangeEmailCode()
			tt.setup(actualChange)

			require.Error(t, actualChange.ValidateData(),
				"expected ValidateData to return an error when %s is empty", tt.name)
		})
	}
}

// ============ ChangePasswordCode tests ============
func TestChangePasswordCode_IsCreatedCorrectly(t *testing.T) {
	actualChange := validChangePasswordCode()

	_, ok := actualChange.GetBodyData().(emailmessage.ChangePasswordCodeBody)
	require.True(t, ok, "expected body data to be of type ChangePasswordCodeBody")

	assert.Equal(t, to, actualChange.To,
		"expected To to match the provided value")
	assert.Equal(t, subject, actualChange.Subject,
		"expected Subject to match the provided value")
	assert.Equal(t, verificationCode, actualChange.VerificationCode,
		"expected VerificationCode to match the provided value")
	assert.Equal(t, codeExpirationTime, actualChange.CodeExpirationTime,
		"expected CodeExpirationTime to match the provided value")
	assert.Equal(t, emailmessage.EmailTypeChangePasswordCode, actualChange.GetEmailType(),
		"expected email type to be ChangePasswordCode")
	assert.Nil(t, actualChange.ValidateData(),
		"expected ValidateData to return nil for a valid ChangePasswordCode")
}

func TestChangePasswordCode_EmptyField_ReturnError(t *testing.T) {
	tests := []struct {
		name  string
		setup func(p *emailmessage.ChangePasswordCode)
	}{
		{
			name: "empty To",
			setup: func(p *emailmessage.ChangePasswordCode) {
				p.To = ""
			},
		},
		{
			name: "empty Subject",
			setup: func(p *emailmessage.ChangePasswordCode) {
				p.Subject = ""
			},
		},
		{
			name: "empty VerificationCode",
			setup: func(p *emailmessage.ChangePasswordCode) {
				p.VerificationCode = ""
			},
		},
		{
			name: "empty CodeExpirationTime",
			setup: func(p *emailmessage.ChangePasswordCode) {
				p.CodeExpirationTime = ""
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualChange := validChangePasswordCode()
			tt.setup(actualChange)

			require.Error(t, actualChange.ValidateData(),
				"expected ValidateData to return an error when %s is empty", tt.name)
		})
	}
}

// ============ Deletion tests ============
func TestDeletionCode_IsCreatedCorrectly(t *testing.T) {
	actualDeletion := validDeletionCode()

	_, ok := actualDeletion.GetBodyData().(emailmessage.DeletionCodeBody)
	require.True(t, ok, "expected body data to be of type DeletionCodeBody")

	assert.Equal(t, to, actualDeletion.To,
		"expected To to match the provided value")
	assert.Equal(t, subject, actualDeletion.Subject,
		"expected Subject to match the provided value")
	assert.Equal(t, verificationCode, actualDeletion.VerificationCode,
		"expected VerificationCode to match the provided value")
	assert.Equal(t, codeExpirationTime, actualDeletion.CodeExpirationTime,
		"expected CodeExpirationTime to match the provided value")
	assert.Equal(t, emailmessage.EmailTypeDeletionCode, actualDeletion.GetEmailType(),
		"expected email type to be DeletionCode")
	assert.Nil(t, actualDeletion.ValidateData(),
		"expected ValidateData to return nil for a valid DeletionCode")
}

func TestDeletionCode_EmptyField_ReturnError(t *testing.T) {
	tests := []struct {
		name  string
		setup func(p *emailmessage.DeletionCode)
	}{
		{
			name: "empty To",
			setup: func(p *emailmessage.DeletionCode) {
				p.To = ""
			},
		},
		{
			name: "empty Subject",
			setup: func(p *emailmessage.DeletionCode) {
				p.Subject = ""
			},
		},
		{
			name: "empty VerificationCode",
			setup: func(p *emailmessage.DeletionCode) {
				p.VerificationCode = ""
			},
		},
		{
			name: "empty CodeExpirationTime",
			setup: func(p *emailmessage.DeletionCode) {
				p.CodeExpirationTime = ""
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualDeletion := validDeletionCode()
			tt.setup(actualDeletion)

			require.Error(t, actualDeletion.ValidateData(),
				"expected ValidateData to return an error when %s is empty", tt.name)
		})
	}
}

// ============ ResetPasswordCode tests ============
func TestResetPasswordCode_IsCreatedCorrectly(t *testing.T) {
	actualReset := validResetPasswordCode()

	_, ok := actualReset.GetBodyData().(emailmessage.ResetPasswordCodeBody)
	require.True(t, ok, "expected body data to be of type ResetPasswordCodeBody")

	assert.Equal(t, to, actualReset.To,
		"expected To to match the provided value")
	assert.Equal(t, subject, actualReset.Subject,
		"expected Subject to match the provided value")
	assert.Equal(t, verificationCode, actualReset.VerificationCode,
		"expected VerificationCode to match the provided value")
	assert.Equal(t, link, actualReset.ResetPasswordLink,
		"expected ResetPasswordLink to match the provided value")
	assert.Equal(t, codeExpirationTime, actualReset.CodeExpirationTime,
		"expected CodeExpirationTime to match the provided value")
	assert.Equal(t, emailmessage.EmailTypeResetPasswordCode, actualReset.GetEmailType(),
		"expected email type to be ResetPasswordCode")
	assert.Nil(t, actualReset.ValidateData(),
		"expected ValidateData to return nil for a valid ResetPasswordCode")
}

func TestResetPasswordCode_EmptyField_ReturnError(t *testing.T) {
	tests := []struct {
		name  string
		setup func(p *emailmessage.ResetPasswordCode)
	}{
		{
			name: "empty To",
			setup: func(p *emailmessage.ResetPasswordCode) {
				p.To = ""
			},
		},
		{
			name: "empty Subject",
			setup: func(p *emailmessage.ResetPasswordCode) {
				p.Subject = ""
			},
		},
		{
			name: "empty VerificationCode",
			setup: func(p *emailmessage.ResetPasswordCode) {
				p.VerificationCode = ""
			},
		},
		{
			name: "empty ResetPasswordLink",
			setup: func(p *emailmessage.ResetPasswordCode) {
				p.ResetPasswordLink = ""
			},
		},
		{
			name: "empty CodeExpirationTime",
			setup: func(p *emailmessage.ResetPasswordCode) {
				p.CodeExpirationTime = ""
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualReset := validResetPasswordCode()
			tt.setup(actualReset)

			require.Error(t, actualReset.ValidateData(),
				"expected ValidateData to return an error when %s is empty", tt.name)
		})
	}
}
