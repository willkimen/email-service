package payload_test

import (
	"emailservice/core/application/send_email/payload"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNotifyDeletionPayloadIsCreatedCorrectly(t *testing.T) {
	actualPayload, err := payload.NewNotifyDeletionPayload(
		to, subject, notifyDeletionTemplatePath,
	)

	assert.Nil(t, err)
	assert.Equal(t, to, actualPayload.To)
	assert.Equal(t, subject, actualPayload.Subject)
	assert.NotEmpty(t, actualPayload.Body)
}
