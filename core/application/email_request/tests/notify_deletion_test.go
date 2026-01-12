package emailrequest_test

import (
	"emailservice/core/application/email_request"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNotifyDeletion_IsCreatedCorrectly(t *testing.T) {
	actualNotify := validNotifyDeletion()

	assert.Equal(t, to, actualNotify.To)
	assert.Equal(t, subject, actualNotify.Subject)
	assert.Equal(t, emailrequest.TemplateNotifyDeletionID, actualNotify.TemplateID())
	assert.Nil(t, actualNotify.ValidateData())
}

func TestNotifyDeletion_EmptyField_ReturnError(t *testing.T) {
	tests := []struct {
		name  string
		setup func(p *emailrequest.NotifyDeletion)
	}{
		{
			name: "empty To",
			setup: func(p *emailrequest.NotifyDeletion) {
				p.To = ""
			},
		},
		{
			name: "empty Subject",
			setup: func(p *emailrequest.NotifyDeletion) {
				p.Subject = ""
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualNotify := validNotifyDeletion()
			tt.setup(actualNotify)

			err := actualNotify.ValidateData()
			assert.NotNil(t, err)
		})
	}
}
