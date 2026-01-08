package payload_test

import (
	"emailservice/core/application/send_email/payload"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeletionCodePayloadIsCreatedCorrectly(t *testing.T) {
	actualPayload, err := payload.NewDeletionCodePayload(
		to, subject, verificationCode, deletionCodeTemplatePath,
		codeExpiratinoHours,
	)

	assert.Nil(t, err)
	assert.Equal(t, to, actualPayload.To)
	assert.Equal(t, subject, actualPayload.Subject)
	assert.Equal(t, verificationCode, actualPayload.VerificationCode)
	assert.Equal(t, codeExpiratinoHours, actualPayload.CodeExpirationHours)
	assert.NotEmpty(t, actualPayload.Body)
	assert.Contains(t, actualPayload.Body, verificationCode)
	assert.Contains(t, actualPayload.Body, strconv.Itoa(codeExpiratinoHours))
}
