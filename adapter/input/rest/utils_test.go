package rest_test

import (
	"emailservice/core/application/email_message"
	"github.com/stretchr/testify/mock"
)

type RequestEmailUseCaseMock struct {
	mock.Mock
}

func (m *RequestEmailUseCaseMock) Request(message emailmessage.EmailMessage) error {
	args := m.Called(message)
	return args.Error(0)
}
