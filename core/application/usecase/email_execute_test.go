package usecase_test

import (
	"emailservice/core/application/usecase"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestExecute_ReturnsError_WhenRendererFails(t *testing.T) {
	renderErr := errors.New("render failed")

	usecase := usecase.ExecuteSendEmailUsecase{
		Renderer: fakeRenderer{
			Err: renderErr,
		},
		Sender: fakeSender{},
		Logger: fakeLogger{},
	}

	err := usecase.ExecuteSend(fakeEmailMessage{})

	require.Error(t, err, "expected Execute to return an error when renderer fails")
	assert.Contains(t, err.Error(), "during rendering",
		"expected error message to indicate failure during rendering")
}

func TestExecute_ReturnsError_WhenSenderFails(t *testing.T) {
	sendErr := errors.New("send failed")

	usecase := usecase.ExecuteSendEmailUsecase{
		Renderer: fakeRenderer{
			Body: "<html>body</html>",
		},
		Sender: fakeSender{
			Err: sendErr,
		},
		Logger: fakeLogger{},
	}

	err := usecase.ExecuteSend(fakeEmailMessage{})

	require.Error(t, err, "expected Execute to return an error when sender fails")
	assert.Contains(t, err.Error(), "during sending",
		"expected error message to indicate failure during sending")
}

func TestExecute_ReturnsNil_WhenRenderAndSendSucceed(t *testing.T) {
	usecase := usecase.ExecuteSendEmailUsecase{
		Renderer: fakeRenderer{
			Body: "<html>body</html>",
		},
		Sender: fakeSender{},
		Logger: fakeLogger{},
	}

	err := usecase.ExecuteSend(fakeEmailMessage{})

	require.NoError(t, err,
		"expected Execute to return nil when render and send succeed")
}
