package emailmessage_test

import (
	"emailservice/core/application/email_message"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBaseMessage_InvalidEmailFormat_ReturnError(t *testing.T) {
	p := validActivationCode()
	p.To = "invalid-email"

	err := p.ValidateData()
	assert.NotNil(t, err)

	var emailErr *emailmessage.EmailInvalidFormatError
	assert.ErrorAs(t, err, &emailErr)
}
