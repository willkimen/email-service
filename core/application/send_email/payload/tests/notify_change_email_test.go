package payload_test

import (
	"emailservice/core/application/send_email/payload"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNotifyChangeEmailPayloadIsCreatedCorrectly(t *testing.T) {
	actualPayload, err := payload.NewNotifyChangeEmailPayload(
		to, subject, link, notifyChangeEmailTemplatePath,
	)

	assert.Nil(t, err)
	assert.Equal(t, to, actualPayload.To)
	assert.Equal(t, subject, actualPayload.Subject)
	assert.Equal(t, link, actualPayload.LoginLink)
	assert.NotEmpty(t, actualPayload.Body)
	assert.Contains(t, actualPayload.Body, link)
}
