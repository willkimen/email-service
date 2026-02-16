package emailmessage_test

import (
	"emailservice/core/application/email_message"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNotifyDeletion_IsCreatedCorrectly(t *testing.T) {
	actualNotify := validNotifyDeletion()

	assert.Equal(t, to, actualNotify.To,
		"expected To to match the provided value")
	assert.Equal(t, subject, actualNotify.Subject,
		"expected Subject to match the provided value")
	assert.Equal(t, emailmessage.EmailTypeNotifyDeletion, actualNotify.GetEmailType(),
		"expected email type to be NotifyDeletion")
	assert.Nil(t, actualNotify.ValidateData(),
		"expected ValidateData to return nil for a valid NotifyDeletion")
}

func TestNotifyDeletion_EmptyField_ReturnError(t *testing.T) {
	tests := []struct {
		name  string
		setup func(p *emailmessage.NotifyDeletion)
	}{
		{
			name: "empty To",
			setup: func(p *emailmessage.NotifyDeletion) {
				p.To = ""
			},
		},
		{
			name: "empty Subject",
			setup: func(p *emailmessage.NotifyDeletion) {
				p.Subject = ""
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualNotify := validNotifyDeletion()
			tt.setup(actualNotify)

			require.Error(t, actualNotify.ValidateData(),
				"expected ValidateData to return an error when %s is empty", tt.name)
		})
	}
}
