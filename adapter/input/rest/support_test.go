package rest_test

import (
	"emailservice/core/application/email_message"
	"log/slog"
	"os"

	"github.com/stretchr/testify/mock"
)

var logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))

type RequestEmailUseCaseMock struct {
	mock.Mock
}

func (m *RequestEmailUseCaseMock) Request(message emailmessage.EmailMessage) error {
	args := m.Called(message)
	return args.Error(0)
}
