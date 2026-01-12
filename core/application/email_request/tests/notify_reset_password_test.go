package emailrequest_test

import (
	"emailservice/core/application/email_request"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNotifyResetPassword_IsCreatedCorrectly(t *testing.T) {
	actualNotify := validNotifyResetPassword()

	assert.Equal(t, to, actualNotify.To)
	assert.Equal(t, subject, actualNotify.Subject)
	assert.Equal(t, link, actualNotify.LoginLink)
	assert.Equal(t, emailrequest.TemplateNotifyResetPasswordID, actualNotify.TemplateID())
	assert.Nil(t, actualNotify.ValidateData())
}

func TestNotifyResetPassword_EmptyField_ReturnError(t *testing.T) {
	tests := []struct {
		name  string
		setup func(p *emailrequest.NotifyResetPassword)
	}{
		{
			name: "empty To",
			setup: func(p *emailrequest.NotifyResetPassword) {
				p.To = ""
			},
		},
		{
			name: "empty Subject",
			setup: func(p *emailrequest.NotifyResetPassword) {
				p.Subject = ""
			},
		},
		{
			name: "empty LoginLink",
			setup: func(p *emailrequest.NotifyResetPassword) {
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
