package emailmessage_test

import (
	"emailservice/core/application/email_message"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNotifyChangeEmail_IsCreatedCorrectly(t *testing.T) {
	actualNotify := validNotifyChangeEmail()
	assert.Equal(t, to, actualNotify.To)
	assert.Equal(t, subject, actualNotify.Subject)
	assert.Equal(t, link, actualNotify.LoginLink)
	assert.Equal(t, emailmessage.EmailTypeNotifyChangeEmail, actualNotify.GetEmailType())
	assert.Nil(t, actualNotify.ValidateData())
}

func TestNotifyChangeEmail_EmptyField_ReturnError(t *testing.T) {
	tests := []struct {
		name  string
		setup func(p *emailmessage.NotifyChangeEmail)
	}{
		{
			name: "empty To",
			setup: func(p *emailmessage.NotifyChangeEmail) {
				p.To = ""
			},
		},
		{
			name: "empty Subject",
			setup: func(p *emailmessage.NotifyChangeEmail) {
				p.Subject = ""
			},
		},
		{
			name: "empty LoginLink",
			setup: func(p *emailmessage.NotifyChangeEmail) {
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
