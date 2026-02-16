package emailmessage_test

import (
	"emailservice/core/application/email_message"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestActivationCode_IsCreatedCorrectly(t *testing.T) {
	actualActivation := validActivationCode()

	assert.Equal(t, to, actualActivation.To,
		"expected To to match the provided value")
	assert.Equal(t, subject, actualActivation.Subject,
		"expected Subject to match the provided value")
	assert.Equal(t, verificationCode, actualActivation.VerificationCode,
		"expected VerificationCode to match the provided value")
	assert.Equal(t, link, actualActivation.ActivationLink,
		"expected ActivationLink to match the provided value")
	assert.Equal(t, codeExpiratinoHours, actualActivation.CodeExpirationHours,
		"expected CodeExpirationHours to match the provided value")
	assert.Equal(t, activationDeadlineDays, actualActivation.ActivationDeadlineDays,
		"expected ActivationDeadlineDays to match the provided value")
	assert.Equal(t, emailmessage.EmailTypeActivationCode, actualActivation.GetEmailType(),
		"expected email type to be ActivationCode")
	assert.Nil(t, actualActivation.ValidateData(),
		"expected ValidateData to return nil for a valid ActivationCode")
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

			require.Error(t, actualActivation.ValidateData(),
				"expected ValidateData to return an error when %s is empty", tt.name)
		})
	}
}
