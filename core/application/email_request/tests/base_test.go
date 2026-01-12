package emailrequest_test

import (
	"emailservice/core/application/email_request"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBase_InvalidEmailFormat_ReturnError(t *testing.T) {
	p := validActivationCode()
	p.To = "invalid-email"

	err := p.ValidateData()
	assert.NotNil(t, err)

	var emailErr *emailrequest.EmailInvalidFormatError
	assert.ErrorAs(t, err, &emailErr)
}
