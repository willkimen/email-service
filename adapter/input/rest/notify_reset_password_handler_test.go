package rest_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"emailservice/adapter/input/rest"
	"emailservice/core/application/email_message"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestNotifyResetPasswordHandler_WhenRequestBodyIsInvalidJSON_ShouldReturnBadRequest(t *testing.T) {
	usecaseMock := new(RequestEmailUseCaseMock)

	handler := rest.HandlerEmail{
		Usecase: usecaseMock,
	}

	req := httptest.NewRequest(
		http.MethodPost,
		"/emails/notify-reset-password",
		strings.NewReader("{invalid-json"),
	)
	w := httptest.NewRecorder()

	handler.NotifyResetPasswordHandler(w, req)

	res := w.Result()

	require.Equal(t, http.StatusBadRequest, res.StatusCode)
	usecaseMock.AssertNotCalled(t, "Request", mock.Anything)
}

func TestNotifyResetPasswordHandler_WhenValidationFails_ShouldReturnUnprocessableEntity(t *testing.T) {
	usecaseMock := new(RequestEmailUseCaseMock)

	usecaseMock.
		On("Request", mock.Anything).
		Return(emailmessage.NewEmptyFieldError("login_link"))

	handler := rest.HandlerEmail{
		Usecase: usecaseMock,
	}

	body := `{
		"to": "user@test.com",
		"subject": "Password reset",
		"login_link": ""
	}`

	req := httptest.NewRequest(
		http.MethodPost,
		"/emails/notify-reset-password",
		strings.NewReader(body),
	)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	handler.NotifyResetPasswordHandler(w, req)

	res := w.Result()

	require.Equal(t, http.StatusUnprocessableEntity, res.StatusCode)
	usecaseMock.AssertCalled(t, "Request", mock.Anything)
}

func TestNotifyResetPasswordHandler_WhenRequestIsValid_ShouldReturnAccepted(t *testing.T) {
	usecaseMock := new(RequestEmailUseCaseMock)

	usecaseMock.
		On("Request", mock.Anything).
		Return(nil)

	handler := rest.HandlerEmail{
		Usecase: usecaseMock,
	}

	body := `{
		"to": "user@test.com",
		"subject": "Password reset",
		"login_link": "https://example.com/login"
	}`

	req := httptest.NewRequest(
		http.MethodPost,
		"/emails/notify-reset-password",
		strings.NewReader(body),
	)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	handler.NotifyResetPasswordHandler(w, req)

	res := w.Result()

	require.Equal(t, http.StatusAccepted, res.StatusCode)
	usecaseMock.AssertCalled(t, "Request", mock.Anything)
}
