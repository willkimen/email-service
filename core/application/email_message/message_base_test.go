package emailmessage_test

import (
	"emailservice/core/application/email_message"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBaseMessage_InvalidEmailFormat_ReturnError(t *testing.T) {
	p := validEmailVerificationCode()
	p.To = "invalid-email"

	err := p.ValidateData()
	require.Error(t, err, "An error is expected due to an invalid email address")

	var emailErr *emailmessage.EmailInvalidFormatError
	require.ErrorAs(t, err, &emailErr,
		"An error of type %T was expected, but %T was delivered", emailErr, err)
}
