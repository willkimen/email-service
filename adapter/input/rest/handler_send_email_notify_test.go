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

// =============== Notify activation tests ===============
func TestNotifyActivationHandler_WhenRequestBodyIsInvalidJSON_ShouldReturnBadRequest(t *testing.T) {
	usecaseMock := new(RequestEmailUseCaseMock)
	handler := rest.NewSendEmailHandler(usecaseMock, logger)

	r := httptest.NewRequest(
		http.MethodPost,
		"/emails/notify-activation",
		strings.NewReader("{invalid-json"),
	)
	w := httptest.NewRecorder()

	handler.NotifyActivationHandler(w, r)
	response := w.Result()

	assert.Equal(t,
		http.StatusBadRequest,
		response.StatusCode,
		"expected status 400 when request body contains invalid JSON",
	)

	usecaseMock.AssertNotCalled(t,
		"Request",
		mock.Anything,
	)
}

func TestNotifyActivationHandler_WhenValidationFails_ShouldReturnUnprocessableEntity(t *testing.T) {
	usecaseMock := new(RequestEmailUseCaseMock)

	usecaseMock.
		On("Request", mock.Anything).
		Return(emailmessage.NewEmptyFieldError("login_link"))

	handler := rest.NewSendEmailHandler(usecaseMock, logger)

	body := `{
		"to": "user@test.com",
		"subject": "Activation",
		"login_link": ""
	}`

	r := httptest.NewRequest(
		http.MethodPost,
		"/emails/notify-activation",
		strings.NewReader(body),
	)
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	handler.NotifyActivationHandler(w, r)
	response := w.Result()

	assert.Equal(t,
		http.StatusUnprocessableEntity,
		response.StatusCode,
		"expected status 422 when validation error occurs",
	)

	usecaseMock.AssertCalled(t,
		"Request",
		mock.Anything,
	)
}

func TestNotifyActivationHandler_WhenRequestIsValid_ShouldReturnAccepted(t *testing.T) {
	usecaseMock := new(RequestEmailUseCaseMock)

	usecaseMock.
		On("Request", mock.Anything).
		Return(nil)

	handler := rest.NewSendEmailHandler(usecaseMock, logger)

	body := `{
		"to": "user@test.com",
		"subject": "Activation",
		"login_link": "https://example.com/login"
	}`

	r := httptest.NewRequest(
		http.MethodPost,
		"/emails/notify-activation",
		strings.NewReader(body),
	)
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	handler.NotifyActivationHandler(w, r)
	response := w.Result()

	assert.Equal(t,
		http.StatusAccepted,
		response.StatusCode,
		"expected status 202 when request is successfully accepted",
	)

	usecaseMock.AssertCalled(t,
		"Request",
		mock.Anything,
	)
}

func TestNotifyActivationHandler_WhenUnexpectedErrorOccurs_ShouldReturnInternalServerError(t *testing.T) {
	usecaseMock := new(RequestEmailUseCaseMock)

	usecaseMock.
		On("Request", mock.Anything).
		Return(errors.New("failed to request email sending"))

	handler := rest.NewSendEmailHandler(usecaseMock, logger)

	body := `{
		"to": "user@test.com",
		"subject": "Account activated",
		"login_link": "https://example.com/login"
	}`

	r := httptest.NewRequest(
		http.MethodPost,
		"/emails/notify-activation",
		strings.NewReader(body),
	)
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	handler.NotifyActivationHandler(w, r)
	response := w.Result()

	assert.Equal(t,
		http.StatusInternalServerError,
		response.StatusCode,
		"expected status 500 when an unexpected error occurs",
	)

	usecaseMock.AssertCalled(t,
		"Request",
		mock.Anything,
	)
}

// =============== Notify change email tests ===============
func TestNotifyChangeEmailHandler_WhenRequestBodyIsInvalidJSON_ShouldReturnBadRequest(t *testing.T) {
	usecaseMock := new(RequestEmailUseCaseMock)
	handler := rest.NewSendEmailHandler(usecaseMock, logger)

	r := httptest.NewRequest(
		http.MethodPost,
		"/emails/notify-change-email",
		strings.NewReader("{invalid-json"),
	)
	w := httptest.NewRecorder()

	handler.NotifyChangeEmailHandler(w, r)
	response := w.Result()

	assert.Equal(
		t,
		http.StatusBadRequest,
		response.StatusCode,
		"expected status 400 when request body contains invalid JSON",
	)

	usecaseMock.AssertNotCalled(t, "Request", mock.Anything)
}

func TestNotifyChangeEmailHandler_WhenValidationFails_ShouldReturnUnprocessableEntity(t *testing.T) {
	usecaseMock := new(RequestEmailUseCaseMock)

	usecaseMock.
		On("Request", mock.Anything).
		Return(emailmessage.NewEmptyFieldError("login_link"))

	handler := rest.NewSendEmailHandler(usecaseMock, logger)

	body := `{
		"to": "user@test.com",
		"subject": "Change email",
		"login_link": ""
	}`

	r := httptest.NewRequest(
		http.MethodPost,
		"/emails/notify-change-email",
		strings.NewReader(body),
	)
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	handler.NotifyChangeEmailHandler(w, r)
	response := w.Result()

	assert.Equal(
		t,
		http.StatusUnprocessableEntity,
		response.StatusCode,
		"expected status 422 when validation error occurs",
	)

	usecaseMock.AssertCalled(t, "Request", mock.Anything)
}

func TestNotifyChangeEmailHandler_WhenRequestIsValid_ShouldReturnAccepted(t *testing.T) {
	usecaseMock := new(RequestEmailUseCaseMock)

	usecaseMock.
		On("Request", mock.Anything).
		Return(nil)

	handler := rest.NewSendEmailHandler(usecaseMock, logger)

	body := `{
		"to": "user@test.com",
		"subject": "Change email",
		"login_link": "https://example.com/login"
	}`

	r := httptest.NewRequest(
		http.MethodPost,
		"/emails/notify-change-email",
		strings.NewReader(body),
	)
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	handler.NotifyChangeEmailHandler(w, r)
	response := w.Result()

	assert.Equal(
		t,
		http.StatusAccepted,
		response.StatusCode,
		"expected status 202 when request is successfully accepted",
	)

	usecaseMock.AssertCalled(t, "Request", mock.Anything)
}

func TestNotifyChangeEmailHandler_WhenUnexpectedErrorOccurs_ShouldReturnInternalServerError(t *testing.T) {
	usecaseMock := new(RequestEmailUseCaseMock)

	usecaseMock.
		On("Request", mock.Anything).
		Return(errors.New("failed to request email sending"))

	handler := rest.NewSendEmailHandler(usecaseMock, logger)

	body := `{
		"to": "user@test.com",
		"subject": "Email changed",
		"login_link": "https://example.com/login"
	}`

	r := httptest.NewRequest(
		http.MethodPost,
		"/emails/notify-change-email",
		strings.NewReader(body),
	)
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	handler.NotifyChangeEmailHandler(w, r)
	response := w.Result()

	assert.Equal(
		t,
		http.StatusInternalServerError,
		response.StatusCode,
		"expected status 500 when an unexpected error occurs",
	)

	usecaseMock.AssertCalled(t, "Request", mock.Anything)
}

// =============== Notify change password tests ===============
func TestNotifyChangePasswordHandler_WhenRequestBodyIsInvalidJSON_ShouldReturnBadRequest(t *testing.T) {
	usecaseMock := new(RequestEmailUseCaseMock)
	handler := rest.NewSendEmailHandler(usecaseMock, logger)

	r := httptest.NewRequest(
		http.MethodPost,
		"/emails/notify-change-password",
		strings.NewReader("{invalid-json"),
	)
	w := httptest.NewRecorder()

	handler.NotifyChangePasswordHandler(w, r)
	response := w.Result()

	assert.Equal(
		t,
		http.StatusBadRequest,
		response.StatusCode,
		"expected status 400 when request body contains invalid JSON",
	)

	usecaseMock.AssertNotCalled(t, "Request", mock.Anything)
}

func TestNotifyChangePasswordHandler_WhenValidationFails_ShouldReturnUnprocessableEntity(t *testing.T) {
	usecaseMock := new(RequestEmailUseCaseMock)

	usecaseMock.
		On("Request", mock.Anything).
		Return(emailmessage.NewEmptyFieldError("login_link"))

	handler := rest.NewSendEmailHandler(usecaseMock, logger)

	body := `{
		"to": "user@test.com",
		"subject": "Password changed",
		"login_link": ""
	}`

	r := httptest.NewRequest(
		http.MethodPost,
		"/emails/notify-change-password",
		strings.NewReader(body),
	)
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	handler.NotifyChangePasswordHandler(w, r)
	response := w.Result()

	assert.Equal(
		t,
		http.StatusUnprocessableEntity,
		response.StatusCode,
		"expected status 422 when validation error occurs",
	)

	usecaseMock.AssertCalled(t, "Request", mock.Anything)
}

func TestNotifyChangePasswordHandler_WhenRequestIsValid_ShouldReturnAccepted(t *testing.T) {
	usecaseMock := new(RequestEmailUseCaseMock)

	usecaseMock.
		On("Request", mock.Anything).
		Return(nil)

	handler := rest.NewSendEmailHandler(usecaseMock, logger)

	body := `{
		"to": "user@test.com",
		"subject": "Password changed",
		"login_link": "https://example.com/login"
	}`

	r := httptest.NewRequest(
		http.MethodPost,
		"/emails/notify-change-password",
		strings.NewReader(body),
	)
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	handler.NotifyChangePasswordHandler(w, r)
	response := w.Result()

	assert.Equal(
		t,
		http.StatusAccepted,
		response.StatusCode,
		"expected status 202 when request is successfully accepted",
	)

	usecaseMock.AssertCalled(t, "Request", mock.Anything)
}

func TestNotifyChangePasswordHandler_WhenUnexpectedErrorOccurs_ShouldReturnInternalServerError(t *testing.T) {
	usecaseMock := new(RequestEmailUseCaseMock)

	usecaseMock.
		On("Request", mock.Anything).
		Return(errors.New("failed to request email sending"))

	handler := rest.NewSendEmailHandler(usecaseMock, logger)

	body := `{
		"to": "user@test.com",
		"subject": "Password changed",
		"login_link": "https://example.com/login"
	}`

	r := httptest.NewRequest(
		http.MethodPost,
		"/emails/notify-change-password",
		strings.NewReader(body),
	)
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	handler.NotifyChangePasswordHandler(w, r)
	response := w.Result()

	assert.Equal(
		t,
		http.StatusInternalServerError,
		response.StatusCode,
		"expected status 500 when an unexpected error occurs",
	)

	usecaseMock.AssertCalled(t, "Request", mock.Anything)
}

// =============== Notify deletion tests ===============
func TestNotifyDeletionHandler_WhenRequestBodyIsInvalidJSON_ShouldReturnBadRequest(t *testing.T) {
	usecaseMock := new(RequestEmailUseCaseMock)
	handler := rest.NewSendEmailHandler(usecaseMock, logger)

	r := httptest.NewRequest(
		http.MethodPost,
		"/emails/notify-deletion",
		strings.NewReader("{invalid-json"),
	)
	w := httptest.NewRecorder()

	handler.NotifyDeletionHandler(w, r)
	response := w.Result()

	assert.Equal(
		t,
		http.StatusBadRequest,
		response.StatusCode,
		"expected status 400 when request body contains invalid JSON",
	)

	usecaseMock.AssertNotCalled(t, "Request", mock.Anything)
}

func TestNotifyDeletionHandler_WhenValidationFails_ShouldReturnUnprocessableEntity(t *testing.T) {
	usecaseMock := new(RequestEmailUseCaseMock)

	usecaseMock.
		On("Request", mock.Anything).
		Return(emailmessage.NewEmptyFieldError("to"))

	handler := rest.NewSendEmailHandler(usecaseMock, logger)

	body := `{
		"to": "",
		"subject": "Account deleted"
	}`

	r := httptest.NewRequest(
		http.MethodPost,
		"/emails/notify-deletion",
		strings.NewReader(body),
	)
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	handler.NotifyDeletionHandler(w, r)
	response := w.Result()

	assert.Equal(
		t,
		http.StatusUnprocessableEntity,
		response.StatusCode,
		"expected status 422 when validation error occurs",
	)

	usecaseMock.AssertCalled(t, "Request", mock.Anything)
}

func TestNotifyDeletionHandler_WhenRequestIsValid_ShouldReturnAccepted(t *testing.T) {
	usecaseMock := new(RequestEmailUseCaseMock)

	usecaseMock.
		On("Request", mock.Anything).
		Return(nil)

	handler := rest.NewSendEmailHandler(usecaseMock, logger)

	body := `{
		"to": "user@test.com",
		"subject": "Account deleted"
	}`

	r := httptest.NewRequest(
		http.MethodPost,
		"/emails/notify-deletion",
		strings.NewReader(body),
	)
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	handler.NotifyDeletionHandler(w, r)
	response := w.Result()

	assert.Equal(
		t,
		http.StatusAccepted,
		response.StatusCode,
		"expected status 202 when request is successfully accepted",
	)

	usecaseMock.AssertCalled(t, "Request", mock.Anything)
}

func TestNotifyDeletionHandler_WhenUnexpectedErrorOccurs_ShouldReturnInternalServerError(t *testing.T) {
	usecaseMock := new(RequestEmailUseCaseMock)

	usecaseMock.
		On("Request", mock.Anything).
		Return(errors.New("failed to request email sending"))

	handler := rest.NewSendEmailHandler(usecaseMock, logger)

	body := `{
		"to": "user@test.com",
		"subject": "Account deleted"
	}`

	r := httptest.NewRequest(
		http.MethodPost,
		"/emails/notify-deletion",
		strings.NewReader(body),
	)
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	handler.NotifyDeletionHandler(w, r)
	response := w.Result()

	assert.Equal(
		t,
		http.StatusInternalServerError,
		response.StatusCode,
		"expected status 500 when an unexpected error occurs",
	)

	usecaseMock.AssertCalled(t, "Request", mock.Anything)
}

// =============== Notify reset password tests ===============
func TestNotifyResetPasswordHandler_WhenRequestBodyIsInvalidJSON_ShouldReturnBadRequest(t *testing.T) {
	usecaseMock := new(RequestEmailUseCaseMock)
	handler := rest.NewSendEmailHandler(usecaseMock, logger)

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

	handler := rest.NewSendEmailHandler(usecaseMock, logger)

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

	handler := rest.NewSendEmailHandler(usecaseMock, logger)

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

	handler := rest.NewSendEmailHandler(usecaseMock, logger)

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
