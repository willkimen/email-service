package emailmessage_test

import (
	"emailservice/core/application/email_message"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

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
	assert.Equal(t, codeExpiratinoHours, actualReset.CodeExpirationHours,
		"expected CodeExpirationHours to match the provided value")
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
			name: "empty CodeExpirationHours",
			setup: func(p *emailmessage.ResetPasswordCode) {
				p.CodeExpirationHours = ""
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
