package rest_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"emailservice/adapter/input/rest"
	"emailservice/core/application/email_message"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNotifyResetPasswordHandler_WhenRequestBodyIsInvalidJSON_ShouldReturnBadRequest(t *testing.T) {
	usecaseMock := new(RequestEmailUseCaseMock)
	handler := rest.NewSendEmailHandler(usecaseMock)

	r := httptest.NewRequest(
		http.MethodPost,
		"/emails/notify-reset-password",
		strings.NewReader("{invalid-json"),
	)
	w := httptest.NewRecorder()

	handler.NotifyResetPasswordHandler(w, r)
	response := w.Result()

	assert.Equal(
		t,
		http.StatusBadRequest,
		response.StatusCode,
		"expected status 400 when request body contains invalid JSON",
	)

	usecaseMock.AssertNotCalled(t, "Request", mock.Anything)
}

func TestNotifyResetPasswordHandler_WhenValidationFails_ShouldReturnUnprocessableEntity(t *testing.T) {
	usecaseMock := new(RequestEmailUseCaseMock)

	usecaseMock.
		On("Request", mock.Anything).
		Return(emailmessage.NewEmptyFieldError("login_link"))

	handler := rest.NewSendEmailHandler(usecaseMock)

	body := `{
		"to": "user@test.com",
		"subject": "Password reset",
		"login_link": ""
	}`

	r := httptest.NewRequest(
		http.MethodPost,
		"/emails/notify-reset-password",
		strings.NewReader(body),
	)
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	handler.NotifyResetPasswordHandler(w, r)
	response := w.Result()

	assert.Equal(
		t,
		http.StatusUnprocessableEntity,
		response.StatusCode,
		"expected status 422 when validation error occurs",
	)

	usecaseMock.AssertCalled(t, "Request", mock.Anything)
}

func TestNotifyResetPasswordHandler_WhenRequestIsValid_ShouldReturnAccepted(t *testing.T) {
	usecaseMock := new(RequestEmailUseCaseMock)

	usecaseMock.
		On("Request", mock.Anything).
		Return(nil)

	handler := rest.NewSendEmailHandler(usecaseMock)

	body := `{
		"to": "user@test.com",
		"subject": "Password reset",
		"login_link": "https://example.com/login"
	}`

	r := httptest.NewRequest(
		http.MethodPost,
		"/emails/notify-reset-password",
		strings.NewReader(body),
	)
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	handler.NotifyResetPasswordHandler(w, r)
	response := w.Result()

	assert.Equal(
		t,
		http.StatusAccepted,
		response.StatusCode,
		"expected status 202 when request is successfully accepted",
	)

	usecaseMock.AssertCalled(t, "Request", mock.Anything)
}

func TestNotifyResetPasswordHandler_WhenUnexpectedErrorOccurs_ShouldReturnInternalServerError(t *testing.T) {
	usecaseMock := new(RequestEmailUseCaseMock)

	usecaseMock.
		On("Request", mock.Anything).
		Return(errors.New("failed to request email sending"))

	handler := rest.NewSendEmailHandler(usecaseMock)

	body := `{
		"to": "user@test.com",
		"subject": "Password reset",
		"login_link": "https://example.com/login"
	}`

	r := httptest.NewRequest(
		http.MethodPost,
		"/emails/notify-reset-password",
		strings.NewReader(body),
	)
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	handler.NotifyResetPasswordHandler(w, r)
	response := w.Result()

	assert.Equal(
		t,
		http.StatusInternalServerError,
		response.StatusCode,
		"expected status 500 when an unexpected error occurs",
	)

	usecaseMock.AssertCalled(t, "Request", mock.Anything)
}
