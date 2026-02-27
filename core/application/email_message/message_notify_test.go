package emailmessage_test

import (
	"emailservice/core/application/email_message"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// ============== NotifyActivation tests ==============
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

// ============== NotifyChangeEmail tests ==============
func TestNotifyChangeEmail_IsCreatedCorrectly(t *testing.T) {
	actualNotify := validNotifyChangeEmail()

	_, ok := actualNotify.GetBodyData().(emailmessage.NotifyChangeEmailBody)
	require.True(t, ok, "expected body data to be of type NotifyChangeEmailBody")

	assert.Equal(t, to, actualNotify.To,
		"expected To to match the provided value")
	assert.Equal(t, subject, actualNotify.Subject,
		"expected Subject to match the provided value")
	assert.Equal(t, link, actualNotify.LoginLink,
		"expected LoginLink to match the provided value")
	assert.Equal(t, emailmessage.EmailTypeNotifyChangeEmail, actualNotify.GetEmailType(),
		"expected email type to be NotifyChangeEmail")
	assert.Nil(t, actualNotify.ValidateData(),
		"expected ValidateData to return nil for a valid NotifyChangeEmail")
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

			require.Error(t, actualNotify.ValidateData(),
				"expected ValidateData to return an error when %s is empty", tt.name)
		})
	}
}

// ============== NotifyChangePassword tests ==============
func TestNotifyChangePasswordEmail_IsCreatedCorrectly(t *testing.T) {
	actualNotify := validNotifyChangePassword()

	_, ok := actualNotify.GetBodyData().(emailmessage.NotifyChangePasswordBody)
	require.True(t, ok, "expected body data to be of type NotifyChangePasswordBody")

	assert.Equal(t, to, actualNotify.To,
		"expected To to match the provided value")
	assert.Equal(t, subject, actualNotify.Subject,
		"expected Subject to match the provided value")
	assert.Equal(t, link, actualNotify.LoginLink,
		"expected LoginLink to match the provided value")
	assert.Equal(t, emailmessage.EmailTypeNotifyChangePassword, actualNotify.GetEmailType(),
		"expected email type to be NotifyChangePassword")
	assert.Nil(t, actualNotify.ValidateData(),
		"expected ValidateData to return nil for a valid NotifyChangePassword")
}

func TestNotifyChangePassword_EmptyField_ReturnError(t *testing.T) {
	tests := []struct {
		name  string
		setup func(p *emailmessage.NotifyChangePassword)
	}{
		{
			name: "empty To",
			setup: func(p *emailmessage.NotifyChangePassword) {
				p.To = ""
			},
		},
		{
			name: "empty Subject",
			setup: func(p *emailmessage.NotifyChangePassword) {
				p.Subject = ""
			},
		},
		{
			name: "empty LoginLink",
			setup: func(p *emailmessage.NotifyChangePassword) {
				p.LoginLink = ""
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualNotify := validNotifyChangePassword()
			tt.setup(actualNotify)

			require.Error(t, actualNotify.ValidateData(),
				"expected ValidateData to return an error when %s is empty", tt.name)
		})
	}
}

// ============== NotifyDeletion tests ==============
func TestNotifyDeletion_IsCreatedCorrectly(t *testing.T) {
	actualNotify := validNotifyDeletion()

	_, ok := actualNotify.GetBodyData().(emailmessage.NotifyDeletionBody)
	require.True(t, ok, "expected body data to be of type NotifyDeletionBody")

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

// ============== NotifyResetPassword tests ==============
func TestNotifyResetPassword_IsCreatedCorrectly(t *testing.T) {
	actualNotify := validNotifyResetPassword()

	_, ok := actualNotify.GetBodyData().(emailmessage.NotifyResetPasswordBody)
	require.True(t, ok, "expected body data to be of type NotifyResetPasswordBody")

	assert.Equal(t, to, actualNotify.To,
		"expected To to match the provided value")
	assert.Equal(t, subject, actualNotify.Subject,
		"expected Subject to match the provided value")
	assert.Equal(t, link, actualNotify.LoginLink,
		"expected LoginLink to match the provided value")
	assert.Equal(t, emailmessage.EmailTypeNotifyResetPassword, actualNotify.GetEmailType(),
		"expected email type to be NotifyResetPassword")
	assert.Nil(t, actualNotify.ValidateData(),
		"expected ValidateData to return nil for a valid NotifyResetPassword")
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

			require.Error(t, actualNotify.ValidateData(),
				"expected ValidateData to return an error when %s is empty", tt.name)
		})
	}
}
