package emailmessage_test

import (
	"emailservice/core/application/email_message"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChangePasswordCode_IsCreatedCorrectly(t *testing.T) {
	actualChange := validChangePasswordCode()

	assert.Equal(t, to, actualChange.To)
	assert.Equal(t, subject, actualChange.Subject)
	assert.Equal(t, verificationCode, actualChange.VerificationCode)
	assert.Equal(t, codeExpiratinoHours, actualChange.CodeExpirationHours)
	assert.Equal(t, emailmessage.EmailTypeChangePasswordCode, actualChange.GetEmailType())
	assert.Nil(t, actualChange.ValidateData())
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

			err := actualChange.ValidateData()
			assert.NotNil(t, err)
		})
	}
}
