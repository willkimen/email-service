package emailrequest_test

import (
	"emailservice/core/application/email_request"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNotifyChangeEmail_IsCreatedCorrectly(t *testing.T) {
	actualNotify := validNotifyChangeEmail()
	assert.Equal(t, to, actualNotify.To)
	assert.Equal(t, subject, actualNotify.Subject)
	assert.Equal(t, link, actualNotify.LoginLink)
	assert.Equal(t, emailrequest.TemplateNotifyChangeEmailID, actualNotify.TemplateID())
	assert.Nil(t, actualNotify.ValidateData())
}

func TestNotifyChangeEmail_EmptyField_ReturnError(t *testing.T) {
	tests := []struct {
		name  string
		setup func(p *emailrequest.NotifyChangeEmail)
	}{
		{
			name: "empty To",
			setup: func(p *emailrequest.NotifyChangeEmail) {
				p.To = ""
			},
		},
		{
			name: "empty Subject",
			setup: func(p *emailrequest.NotifyChangeEmail) {
				p.Subject = ""
			},
		},
		{
			name: "empty LoginLink",
			setup: func(p *emailrequest.NotifyChangeEmail) {
				p.LoginLink = ""
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualNotify := validNotifyChangeEmail()
			tt.setup(actualNotify)

			err := actualNotify.ValidateData()
			assert.NotNil(t, err)
		})
	}
}
