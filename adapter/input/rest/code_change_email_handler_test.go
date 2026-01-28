package rest_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"emailservice/adapter/input/rest"
	"emailservice/core/application/email_message"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestSendChangeEmailCodeHandler_WhenRequestBodyIsInvalidJSON_ShouldReturnBadRequest(t *testing.T) {
	usecaseMock := new(RequestEmailUseCaseMock)

	handler := rest.HandlerEmail{
		Usecase: usecaseMock,
	}

	req := httptest.NewRequest(
		http.MethodPost,
		"/emails/change-email-code",
		strings.NewReader("{invalid-json"),
	)
	w := httptest.NewRecorder()

	handler.SendChangeEmailCodeHandler(w, req)

	res := w.Result()

	require.Equal(t, http.StatusBadRequest, res.StatusCode)
	usecaseMock.AssertNotCalled(t, "Request", mock.Anything)
}

func TestSendChangeEmailCodeHandler_WhenValidationFails_ShouldReturnUnprocessableEntity(t *testing.T) {
	usecaseMock := new(RequestEmailUseCaseMock)

	usecaseMock.
		On("Request", mock.Anything).
		Return(emailmessage.NewEmptyFieldError("verification_code"))

	handler := rest.HandlerEmail{
		Usecase: usecaseMock,
	}

	body := `{
		"to": "user@test.com",
		"subject": "Change email",
		"verification_code": "",
		"code_expiration_hours": "2"
	}`

	req := httptest.NewRequest(
		http.MethodPost,
		"/emails/change-email-code",
		strings.NewReader(body),
	)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	handler.SendChangeEmailCodeHandler(w, req)

	res := w.Result()

	require.Equal(t, http.StatusUnprocessableEntity, res.StatusCode)
	usecaseMock.AssertCalled(t, "Request", mock.Anything)
}

func TestSendChangeEmailCodeHandler_WhenRequestIsValid_ShouldReturnAccepted(t *testing.T) {
	usecaseMock := new(RequestEmailUseCaseMock)

	usecaseMock.
		On("Request", mock.Anything).
		Return(nil)

	handler := rest.HandlerEmail{
		Usecase: usecaseMock,
	}

	body := `{
		"to": "user@test.com",
		"subject": "Change email",
		"verification_code": "123456",
		"code_expiration_hours": "2"
	}`

	req := httptest.NewRequest(
		http.MethodPost,
		"/emails/change-email-code",
		strings.NewReader(body),
	)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	handler.SendChangeEmailCodeHandler(w, req)

	res := w.Result()

	require.Equal(t, http.StatusAccepted, res.StatusCode)
	usecaseMock.AssertCalled(t, "Request", mock.Anything)
}

func TestSendChangeEmailCodeHandler_WhenUnexpectedErrorOccurs_ShouldReturnInternalServerError(t *testing.T) {
	usecaseMock := new(RequestEmailUseCaseMock)

	usecaseMock.
		On("Request", mock.Anything).
		Return(errors.New("failed to request email sending"))

	handler := rest.HandlerEmail{
		Usecase: usecaseMock,
	}

	body := `{
		"to": "user@test.com",
		"subject": "Change email",
		"verification_code": "123456",
		"code_expiration_hours": "2"
	}`

	req := httptest.NewRequest(
		http.MethodPost,
		"/emails/change-email-code",
		strings.NewReader(body),
	)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	handler.SendChangeEmailCodeHandler(w, req)

	res := w.Result()

	require.Equal(t, http.StatusInternalServerError, res.StatusCode)
	usecaseMock.AssertCalled(t, "Request", mock.Anything)
}
