package emailrequest_test

import (
	"emailservice/core/application/email_request"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeletionCode_IsCreatedCorrectly(t *testing.T) {
	actualDeletion := validDeletionCode()

	assert.Equal(t, to, actualDeletion.To)
	assert.Equal(t, subject, actualDeletion.Subject)
	assert.Equal(t, verificationCode, actualDeletion.VerificationCode)
	assert.Equal(t, codeExpiratinoHours, actualDeletion.CodeExpirationHours)
	assert.Equal(t, emailrequest.TemplateDeletionCodeID, actualDeletion.TemplateID())
	assert.Nil(t, actualDeletion.ValidateData())
}

func TestDeletionCode_EmptyField_ReturnError(t *testing.T) {
	tests := []struct {
		name  string
		setup func(p *emailrequest.DeletionCode)
	}{
		{
			name: "empty To",
			setup: func(p *emailrequest.DeletionCode) {
				p.To = ""
			},
		},
		{
			name: "empty Subject",
			setup: func(p *emailrequest.DeletionCode) {
				p.Subject = ""
			},
		},
		{
			name: "empty VerificationCode",
			setup: func(p *emailrequest.DeletionCode) {
				p.VerificationCode = ""
			},
		},
		{
			name: "empty CodeExpirationHours",
			setup: func(p *emailrequest.DeletionCode) {
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
