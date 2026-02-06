package usecase_test

import (
	"emailservice/core/application/usecase"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExecute_ReturnsError_WhenRendererFails(t *testing.T) {
	renderErr := errors.New("render failed")

	usecase := usecase.ExecuteSendEmailUsecase{
		Renderer: FakeRenderer{
			Err: renderErr,
		},
		Sender: FakeSender{},
	}

	err := usecase.Execute(FakeEmailMessage{})

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "during rendering")
}

func TestExecute_ReturnsError_WhenSenderFails(t *testing.T) {
	sendErr := errors.New("send failed")

	usecase := usecase.ExecuteSendEmailUsecase{
		Renderer: FakeRenderer{
			Body: "<html>body</html>",
		},
		Sender: FakeSender{
			Err: sendErr,
		},
	}

	err := usecase.Execute(FakeEmailMessage{})

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "during sending")
}

func TestExecute_ReturnsNil_WhenRenderAndSendSucceed(t *testing.T) {
	usecase := usecase.ExecuteSendEmailUsecase{
		Renderer: FakeRenderer{
			Body: "<html>body</html>",
		},
		Sender: FakeSender{},
	}

	err := usecase.Execute(FakeEmailMessage{})

	assert.NoError(t, err)
}
