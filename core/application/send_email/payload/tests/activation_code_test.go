package payload_test

import (
	"emailservice/core/application/send_email/payload"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestActivationCodePayloadIsCreatedCorrectly(t *testing.T) {
	actualPayload, err := payload.NewActivationCodePayload(
		to, subject, verificationCode, link,
		codeExpiratinoHours, activationDeadlineDays,
		activationCodeTemplatePath,
	)

	assert.Nil(t, err)
	assert.Equal(t, to, actualPayload.To)
	assert.Equal(t, subject, actualPayload.Subject)
	assert.Equal(t, verificationCode, actualPayload.VerificationCode)
	assert.Equal(t, link, actualPayload.ActivationLink)
	assert.Equal(t, codeExpiratinoHours, actualPayload.CodeExpirationHours)
	assert.Equal(t, activationDeadlineDays, actualPayload.ActivationDeadlineDays)
	assert.NotEmpty(t, actualPayload.Body)
	assert.Contains(t, actualPayload.Body, verificationCode)
	assert.Contains(t, actualPayload.Body, link)
	assert.Contains(t, actualPayload.Body, strconv.Itoa(codeExpiratinoHours))
	assert.Contains(t, actualPayload.Body, strconv.Itoa(activationDeadlineDays))
}
