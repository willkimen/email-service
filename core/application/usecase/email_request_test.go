package usecase_test

import (
	"emailservice/core/application/usecase"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestExecute_Success(t *testing.T) {
	usecase := usecase.RequestSendEmailUseCase{
		Publisher: fakePublisher{
			Err: nil,
		},
	}

	err := usecase.Execute(*messageDummy)

	assert.Nil(t, err)
}

func TestExecute_ReturnError_RequestFails(t *testing.T) {
	usecase := usecase.RequestSendEmailUseCase{
		Publisher: fakePublisher{
			Err: originalCause,
		},
	}

	err := usecase.Execute(*messageDummy)

	require.Error(t, err)
}
