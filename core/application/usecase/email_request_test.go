package usecase_test

import (
	"emailservice/core/application/email_message"
	"emailservice/core/application/usecase"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWhenMessageIsInvalid_RequestIsRejected(t *testing.T) {
	request := usecase.RequestSendEmailUseCase{Publisher: publisherSuccessFake{}}
	err := request.Request(invalidMessage)

	assert.NotNil(t, err)
	assert.IsType(t, &emailmessage.EmptyFieldError{}, err)

	var fvErr emailmessage.FieldValidationError
	assert.ErrorAs(t, err, &fvErr)
}

func TestWhenPublisherFails_RequestFails(t *testing.T) {
	request := usecase.RequestSendEmailUseCase{Publisher: publishFailureFake{}}
	err := request.Request(messageCorrect)

	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "failed to request email sending:")
}

func TestSuccessfulRequest(t *testing.T) {
	request := usecase.RequestSendEmailUseCase{Publisher: publisherSuccessFake{}}
	err := request.Request(messageCorrect)

	assert.Nil(t, err)
}
