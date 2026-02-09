package emailmessage_test

import (
	"emailservice/core/application/email_message"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNotifyResetPassword_IsCreatedCorrectly(t *testing.T) {
	actualNotify := validNotifyResetPassword()

	assert.Equal(t, to, actualNotify.To)
	assert.Equal(t, subject, actualNotify.Subject)
	assert.Equal(t, link, actualNotify.LoginLink)
	assert.Equal(t, emailmessage.EmailTypeNotifyResetPassword, actualNotify.GetEmailType())
	assert.Nil(t, actualNotify.ValidateData())
}

func TestNotifyResetPassword_EmptyField_ReturnError(t *testing.T) {
	tests := []struct {
		name  string
		setup func(p *emailmessage.NotifyResetPassword)
	}{
		{
			name: "empty To",
			setup: func(p *emailmessage.NotifyResetPassword) {
				p.To = ""
			},
		},
		{
			name: "empty Subject",
			setup: func(p *emailmessage.NotifyResetPassword) {
				p.Subject = ""
			},
		},
		{
			name: "empty LoginLink",
			setup: func(p *emailmessage.NotifyResetPassword) {
				p.LoginLink = ""
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualNotify := validNotifyResetPassword()
			tt.setup(actualNotify)

			err := actualNotify.ValidateData()
			assert.NotNil(t, err)
		})
	}
}
