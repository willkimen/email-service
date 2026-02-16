package emailmessage_test

import (
	"emailservice/core/application/email_message"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDeletionCode_IsCreatedCorrectly(t *testing.T) {
	actualDeletion := validDeletionCode()

	assert.Equal(t, to, actualDeletion.To,
		"expected To to match the provided value")
	assert.Equal(t, subject, actualDeletion.Subject,
		"expected Subject to match the provided value")
	assert.Equal(t, verificationCode, actualDeletion.VerificationCode,
		"expected VerificationCode to match the provided value")
	assert.Equal(t, codeExpiratinoHours, actualDeletion.CodeExpirationHours,
		"expected CodeExpirationHours to match the provided value")
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
			name: "empty CodeExpirationHours",
			setup: func(p *emailmessage.DeletionCode) {
				p.CodeExpirationHours = ""
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
