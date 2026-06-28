package usecase_test

import (
	"emailservice/core/application/usecase"
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExecute_ReturnsNil_WhenRenderAndSendSucceed(t *testing.T) {
	usecase := usecase.ExecuteSendEmailUsecase{
		Renderer: fakeRenderer{
			Subject: "fake subject",
			Body:    "<html>body</html>",
			Err:     nil,
		},
		Sender: fakeSender{
			Err: nil,
		},
	}

	err := usecase.Execute(*messageDummy)

	require.NoError(t, err)
}

func TestExecute_ReturnsError_WhenRendererFails(t *testing.T) {
	usecase := usecase.ExecuteSendEmailUsecase{
		Renderer: fakeRenderer{
			Subject: "",
			Body:    "",
			Err:     errors.New("render failed"),
		},
		Sender: fakeSender{
			Err: nil,
		},
	}

	err := usecase.Execute(*messageDummy)

	require.Error(t, err)
}

func TestExecute_ReturnsError_WhenSenderFails(t *testing.T) {
	usecase := usecase.ExecuteSendEmailUsecase{
		Renderer: fakeRenderer{
			Subject: "fake subject",
			Body:    "<html>body</html>",
			Err:     nil,
		},
		Sender: fakeSender{
			Err: errors.New("send failed"),
		},
	}

	err := usecase.Execute(*messageDummy)

	require.Error(t, err)
}
