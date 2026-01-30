package emailmessage_test

import (
	"emailservice/core/application/email_message"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestActivationCode_IsCreatedCorrectly(t *testing.T) {
	actualActivation := validActivationCode()

	assert.Equal(t, to, actualActivation.To)
	assert.Equal(t, subject, actualActivation.Subject)
	assert.Equal(t, verificationCode, actualActivation.VerificationCode)
	assert.Equal(t, link, actualActivation.ActivationLink)
	assert.Equal(t, codeExpiratinoHours, actualActivation.CodeExpirationHours)
	assert.Equal(t, activationDeadlineDays, actualActivation.ActivationDeadlineDays)
	assert.Equal(t, emailmessage.TemplateActivationCodeID, actualActivation.TemplateID())
	assert.Nil(t, actualActivation.ValidateData())
}

func TestActivationCode_EmptyField_ReturnError(t *testing.T) {
	tests := []struct {
		name  string
		setup func(p *emailmessage.ActivationCode)
	}{
		{
			name: "empty To",
			setup: func(p *emailmessage.ActivationCode) {
				p.To = ""
			},
		},
		{
			name: "empty Subject",
			setup: func(p *emailmessage.ActivationCode) {
				p.Subject = ""
			},
		},
		{
			name: "empty VerificationCode",
			setup: func(p *emailmessage.ActivationCode) {
				p.VerificationCode = ""
			},
		},
		{
			name: "empty CodeExpirationHours",
			setup: func(p *emailmessage.ActivationCode) {
				p.CodeExpirationHours = ""
			},
		},
		{
			name: "empty ActivationLink",
			setup: func(p *emailmessage.ActivationCode) {
				p.ActivationLink = ""
			},
		},
		{
			name: "empty ActivationDeadlineDays",
			setup: func(p *emailmessage.ActivationCode) {
				p.ActivationDeadlineDays = ""
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualActivation := validActivationCode()
			tt.setup(actualActivation)

			err := actualActivation.ValidateData()
			assert.NotNil(t, err)
		})
	}
}
