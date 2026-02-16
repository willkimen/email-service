package usecase_test

import (
	"emailservice/core/application/email_message"
	"emailservice/core/application/usecase"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWhenMessageIsInvalid_RequestIsRejected(t *testing.T) {
	request := usecase.RequestSendEmailUseCase{
		Publisher: publisherSuccessFake{},
	}

	err := request.Request(invalidMessage)

	require.Error(t, err,
		"expected Request to return an error when message is invalid")

	assert.IsType(t, &emailmessage.EmptyFieldError{}, err,
		"expected error to be of type EmptyFieldError")

	var fvErr emailmessage.FieldValidationError
	assert.ErrorAs(t, err, &fvErr,
		"expected error to implement FieldValidationError")
}

func TestWhenPublisherFails_RequestFails(t *testing.T) {
	request := usecase.RequestSendEmailUseCase{Publisher: publishFailureFake{}}
	err := request.Request(messageCorrect)

	require.Error(t, err,
		"expected Request to return an error when publisher fails")

	assert.Contains(t, err.Error(),
		"failed to request email sending:",
		"expected error message to indicate publisher failure")
}

func TestSuccessfulRequest(t *testing.T) {
	request := usecase.RequestSendEmailUseCase{Publisher: publisherSuccessFake{}}
	err := request.Request(messageCorrect)

	assert.NoError(t, err,
		"expected Request to return nil when message is valid and publisher succeeds")
}
