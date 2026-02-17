package emailmessage_test

import (
	"emailservice/core/application/email_message"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNotifyActivation_IsCreatedCorrectly(t *testing.T) {
	actualNotify := validNotifyActivation()

	_, ok := actualNotify.GetBodyData().(emailmessage.NotifyActivationBody)
	require.True(t, ok, "expected body data to be of type NotifyActivationBody")

	assert.Equal(t, to, actualNotify.To,
		"expected To to match the provided value")
	assert.Equal(t, subject, actualNotify.Subject,
		"expected Subject to match the provided value")
	assert.Equal(t, link, actualNotify.LoginLink,
		"expected LoginLink to match the provided value")
	assert.Equal(t, emailmessage.EmailTypeNotifyActivation, actualNotify.GetEmailType(),
		"expected email type to be NotifyActivation")
	assert.Nil(t, actualNotify.ValidateData(),
		"expected ValidateData to return nil for a valid NotifyActivation")
}

func TestNotifyActivation_EmptyField_ReturnError(t *testing.T) {
	tests := []struct {
		name  string
		setup func(p *emailmessage.NotifyActivation)
	}{
		{
			name: "empty To",
			setup: func(p *emailmessage.NotifyActivation) {
				p.To = ""
			},
		},
		{
			name: "empty Subject",
			setup: func(p *emailmessage.NotifyActivation) {
				p.Subject = ""
			},
		},
		{
			name: "empty LoginLink",
			setup: func(p *emailmessage.NotifyActivation) {
				p.LoginLink = ""
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualNotify := validNotifyActivation()
			tt.setup(actualNotify)

			require.Error(t, actualNotify.ValidateData(),
				"expected ValidateData to return an error when %s is empty", tt.name)
		})
	}
}
