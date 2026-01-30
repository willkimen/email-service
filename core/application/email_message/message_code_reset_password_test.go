package emailmessage_test

import (
	"emailservice/core/application/email_message"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResetPasswordCode_IsCreatedCorrectly(t *testing.T) {
	actualReset := validResetPasswordCode()

	assert.Equal(t, to, actualReset.To)
	assert.Equal(t, subject, actualReset.Subject)
	assert.Equal(t, verificationCode, actualReset.VerificationCode)
	assert.Equal(t, link, actualReset.ResetPasswordLink)
	assert.Equal(t, codeExpiratinoHours, actualReset.CodeExpirationHours)
	assert.Equal(t, emailmessage.TemplateResetPasswordCodeID, actualReset.TemplateID())
	assert.Nil(t, actualReset.ValidateData())
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

			err := actualReset.ValidateData()
			assert.NotNil(t, err)
		})
	}
}
