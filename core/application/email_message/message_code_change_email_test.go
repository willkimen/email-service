package emailmessage_test

import (
	"emailservice/core/application/email_message"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

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
	assert.Equal(t, codeExpiratinoHours, actualChange.CodeExpirationHours,
		"expected CodeExpirationHours to match the provided value")
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
			name: "empty CodeExpirationHours",
			setup: func(p *emailmessage.ChangeEmailCode) {
				p.CodeExpirationHours = ""
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
