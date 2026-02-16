package emailmessage_test

import (
	"emailservice/core/application/email_message"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestChangePasswordCode_IsCreatedCorrectly(t *testing.T) {
	actualChange := validChangePasswordCode()

	assert.Equal(t, to, actualChange.To,
		"expected To to match the provided value")
	assert.Equal(t, subject, actualChange.Subject,
		"expected Subject to match the provided value")
	assert.Equal(t, verificationCode, actualChange.VerificationCode,
		"expected VerificationCode to match the provided value")
	assert.Equal(t, codeExpiratinoHours, actualChange.CodeExpirationHours,
		"expected CodeExpirationHours to match the provided value")
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
			name: "empty CodeExpirationHours",
			setup: func(p *emailmessage.ChangePasswordCode) {
				p.CodeExpirationHours = ""
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
