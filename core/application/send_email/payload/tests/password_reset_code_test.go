package payload_test

import (
	"emailservice/core/application/send_email/payload"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPasswordResetCodePayloadIsCreatedCorrectly(t *testing.T) {
	actualPayload, err := payload.NewPasswordResetCodePayload(
		to, subject, verificationCode, link, passwordResetCodeTemplatePath,
		codeExpiratinoHours,
	)

	assert.Nil(t, err)
	assert.Equal(t, to, actualPayload.To)
	assert.Equal(t, subject, actualPayload.Subject)
	assert.Equal(t, verificationCode, actualPayload.VerificationCode)
	assert.Equal(t, link, actualPayload.ResetPasswordLink)
	assert.Equal(t, codeExpiratinoHours, actualPayload.CodeExpirationHours)
	assert.NotEmpty(t, actualPayload.Body)
	assert.Contains(t, actualPayload.Body, verificationCode)
	assert.Contains(t, actualPayload.Body, link)
	assert.Contains(t, actualPayload.Body, strconv.Itoa(codeExpiratinoHours))
}
