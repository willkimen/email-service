package emailmessage_test

import (
	"emailservice/core/application/email_message"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeletionCode_IsCreatedCorrectly(t *testing.T) {
	actualDeletion := validDeletionCode()

	assert.Equal(t, to, actualDeletion.To)
	assert.Equal(t, subject, actualDeletion.Subject)
	assert.Equal(t, verificationCode, actualDeletion.VerificationCode)
	assert.Equal(t, codeExpiratinoHours, actualDeletion.CodeExpirationHours)
	assert.Equal(t, emailmessage.TemplateDeletionCodeID, actualDeletion.TemplateID())
	assert.Nil(t, actualDeletion.ValidateData())
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
			name: "empty CodeExpirationHours",
			setup: func(p *emailmessage.DeletionCode) {
				p.CodeExpirationHours = ""
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualDeletin := validDeletionCode()
			tt.setup(actualDeletin)

			err := actualDeletin.ValidateData()
			assert.NotNil(t, err)
		})
	}
}
