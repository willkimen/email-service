package payload_test

import (
	"emailservice/core/application/send_email/payload"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNotifyPasswordResetPayloadIsCreatedCorrectly(t *testing.T) {
	actualPayload, err := payload.NewNotifyResetPasswordPayload(
		to, subject, link, notifyPasswordResetTemplatePath,
	)

	assert.Nil(t, err)
	assert.Equal(t, to, actualPayload.To)
	assert.Equal(t, subject, actualPayload.Subject)
	assert.Equal(t, link, actualPayload.LoginLink)
	assert.NotEmpty(t, actualPayload.Body)
	assert.Contains(t, actualPayload.Body, link)
}
