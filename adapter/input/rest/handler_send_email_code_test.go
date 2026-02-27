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

func assertBadRequest(t *testing.T, response *http.Response, usecaseMock *RequestEmailUseCaseMock) {
	t.Helper()

	assert.Equal(t,
		http.StatusBadRequest,
		response.StatusCode,
		"expected status code to be 400 when request body contains invalid JSON",
	)

	usecaseMock.AssertNotCalled(t, "Request", mock.Anything)
}

func assertUnprocessableEntity(t *testing.T, response *http.Response, usecaseMock *RequestEmailUseCaseMock) {
	t.Helper()

	assert.Equal(t,
		http.StatusUnprocessableEntity,
		response.StatusCode,
		"expected status code to be 422 when validation fails",
	)

	usecaseMock.AssertCalled(t, "Request", mock.Anything)
}
func assertAccepted(t *testing.T, response *http.Response, usecaseMock *RequestEmailUseCaseMock) {
	t.Helper()

	assert.Equal(t,
		http.StatusAccepted,
		response.StatusCode,
		"expected status code to be 202 when request is valid",
	)

	usecaseMock.AssertCalled(t, "Request", mock.Anything)
}

func assertInternalServerError(t *testing.T, response *http.Response, usecaseMock *RequestEmailUseCaseMock) {
	t.Helper()

	assert.Equal(t,
		http.StatusInternalServerError,
		response.StatusCode,
		"expected status code to be 500 when unexpected error occurs",
	)

	usecaseMock.AssertCalled(t, "Request", mock.Anything)
}

// =============== Activation code tests ===============
func TestSendActivationCodeHandler_WhenRequestBodyIsInvalidJSON_ShouldReturnBadRequest(t *testing.T) {
	usecaseMock := new(RequestEmailUseCaseMock)

	handler := rest.NewSendEmailHandler(usecaseMock, logger)

	r := httptest.NewRequest(
		http.MethodPost,
		"/emails/activation-code",
		strings.NewReader("{invalid-json"),
	)
	w := httptest.NewRecorder()

	handler.SendActivationCodeHandler(w, r)

	assertBadRequest(t, w.Result(), usecaseMock)
}

func TestSendActivationCodeHandler_WhenValidationFails_ShouldReturnUnprocessableEntity(t *testing.T) {
	usecaseMock := new(RequestEmailUseCaseMock)

	usecaseMock.
		On("Request", mock.Anything).
		Return(emailmessage.NewEmptyFieldError("to"))

	handler := rest.NewSendEmailHandler(usecaseMock, logger)

	body := `{
		"to": "",
		"subject": "Activation",
		"verification_code": "123456",
		"code_expiration_hours": "2",
		"activation_link": "https://example.com",
		"activation_deadline_days": "7"
	}`

	r := httptest.NewRequest(http.MethodPost, "/emails/activation-code", strings.NewReader(body))
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	handler.SendActivationCodeHandler(w, r)

	assertUnprocessableEntity(t, w.Result(), usecaseMock)
}

func TestSendActivationCodeHandler_WhenRequestIsValid_ShouldReturnAccepted(t *testing.T) {
	usecaseMock := new(RequestEmailUseCaseMock)

	usecaseMock.
		On("Request", mock.Anything).
		Return(nil)

	handler := rest.NewSendEmailHandler(usecaseMock, logger)

	body := `{
		"to": "user@test.com",
		"subject": "Activation",
		"verification_code": "123456",
		"code_expiration_hours": "2",
		"activation_link": "https://example.com",
		"activation_deadline_days": "7"
	}`

	r := httptest.NewRequest(http.MethodPost, "/emails/activation-code", strings.NewReader(body))
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	handler.SendActivationCodeHandler(w, r)

	assertAccepted(t, w.Result(), usecaseMock)
}

func TestSendActivationCodeHandler_WhenUnexpectedErrorOccurs_ShouldReturnInternalServerError(t *testing.T) {
	usecaseMock := new(RequestEmailUseCaseMock)

	usecaseMock.
		On("Request", mock.Anything).
		Return(errors.New("failed to request email sending"))

	handler := rest.NewSendEmailHandler(usecaseMock, logger)

	body := `{
		"to": "user@test.com",
		"subject": "Activation",
		"verification_code": "123456",
		"code_expiration_hours": "2",
		"activation_link": "https://example.com",
		"activation_deadline_days": "7"
	}`

	r := httptest.NewRequest(http.MethodPost, "/emails/activation-code", strings.NewReader(body))
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	handler.SendActivationCodeHandler(w, r)

	assertInternalServerError(t, w.Result(), usecaseMock)
}

// =============== Change email code tests ===============
func TestSendChangeEmailCodeHandler_WhenRequestBodyIsInvalidJSON_ShouldReturnBadRequest(t *testing.T) {
	usecaseMock := new(RequestEmailUseCaseMock)

	handler := rest.NewSendEmailHandler(usecaseMock, logger)

	r := httptest.NewRequest(
		http.MethodPost,
		"/emails/change-email-code",
		strings.NewReader("{invalid-json"),
	)
	w := httptest.NewRecorder()

	handler.SendChangeEmailCodeHandler(w, r)

	assertBadRequest(t, w.Result(), usecaseMock)
}

func TestSendChangeEmailCodeHandler_WhenValidationFails_ShouldReturnUnprocessableEntity(t *testing.T) {
	usecaseMock := new(RequestEmailUseCaseMock)

	usecaseMock.
		On("Request", mock.Anything).
		Return(emailmessage.NewEmptyFieldError("verification_code"))

	handler := rest.NewSendEmailHandler(usecaseMock, logger)

	body := `{
		"to": "user@test.com",
		"subject": "Change email",
		"verification_code": "",
		"code_expiration_hours": "2"
	}`

	r := httptest.NewRequest(
		http.MethodPost,
		"/emails/change-email-code",
		strings.NewReader(body),
	)
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	handler.SendChangeEmailCodeHandler(w, r)

	assertUnprocessableEntity(t, w.Result(), usecaseMock)
}

func TestSendChangeEmailCodeHandler_WhenRequestIsValid_ShouldReturnAccepted(t *testing.T) {
	usecaseMock := new(RequestEmailUseCaseMock)

	usecaseMock.
		On("Request", mock.Anything).
		Return(nil)

	handler := rest.NewSendEmailHandler(usecaseMock, logger)

	body := `{
		"to": "user@test.com",
		"subject": "Change email",
		"verification_code": "123456",
		"code_expiration_hours": "2"
	}`

	r := httptest.NewRequest(
		http.MethodPost,
		"/emails/change-email-code",
		strings.NewReader(body),
	)
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	handler.SendChangeEmailCodeHandler(w, r)

	assertAccepted(t, w.Result(), usecaseMock)
}

func TestSendChangeEmailCodeHandler_WhenUnexpectedErrorOccurs_ShouldReturnInternalServerError(t *testing.T) {
	usecaseMock := new(RequestEmailUseCaseMock)

	usecaseMock.
		On("Request", mock.Anything).
		Return(errors.New("failed to request email sending"))

	handler := rest.NewSendEmailHandler(usecaseMock, logger)

	body := `{
		"to": "user@test.com",
		"subject": "Change email",
		"verification_code": "123456",
		"code_expiration_hours": "2"
	}`

	r := httptest.NewRequest(
		http.MethodPost,
		"/emails/change-email-code",
		strings.NewReader(body),
	)
	r.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	handler.SendChangeEmailCodeHandler(w, r)

	assertInternalServerError(t, w.Result(), usecaseMock)
}

// =============== Change password code tests ===============
func TestSendChangePasswordCodeHandler_WhenRequestBodyIsInvalidJSON_ShouldReturnBadRequest(t *testing.T) {
	usecaseMock := new(RequestEmailUseCaseMock)

	handler := rest.NewSendEmailHandler(usecaseMock, logger)

	r := httptest.NewRequest(
		http.MethodPost,
		"/emails/change-password-code",
		strings.NewReader("{invalid-json"),
	)
	w := httptest.NewRecorder()

	handler.SendChangePasswordCodeHandler(w, r)

	assertBadRequest(t, w.Result(), usecaseMock)
}

func TestSendChangePasswordCodeHandler_WhenValidationFails_ShouldReturnUnprocessableEntity(t *testing.T) {
	usecaseMock := new(RequestEmailUseCaseMock)

	usecaseMock.
		On("Request", mock.Anything).
		Return(emailmessage.NewEmptyFieldError("verification_code"))

	handler := rest.NewSendEmailHandler(usecaseMock, logger)

	body := `{
		"to": "user@test.com",
		"subject": "Change password",
		"verification_code": "",
		"code_expiration_hours": "2"
	}`

	r := httptest.NewRequest(
		http.MethodPost,
		"/emails/change-password-code",
		strings.NewReader(body),
	)
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	handler.SendChangePasswordCodeHandler(w, r)

	assertUnprocessableEntity(t, w.Result(), usecaseMock)
}

func TestSendChangePasswordCodeHandler_WhenRequestIsValid_ShouldReturnAccepted(t *testing.T) {
	usecaseMock := new(RequestEmailUseCaseMock)

	usecaseMock.
		On("Request", mock.Anything).
		Return(nil)

	handler := rest.NewSendEmailHandler(usecaseMock, logger)

	body := `{
		"to": "user@test.com",
		"subject": "Change password",
		"verification_code": "123456",
		"code_expiration_hours": "2"
	}`

	r := httptest.NewRequest(
		http.MethodPost,
		"/emails/change-password-code",
		strings.NewReader(body),
	)
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	handler.SendChangePasswordCodeHandler(w, r)

	assertAccepted(t, w.Result(), usecaseMock)
}

func TestSendChangePasswordCodeHandler_WhenUnexpectedErrorOccurs_ShouldReturnInternalServerError(t *testing.T) {
	usecaseMock := new(RequestEmailUseCaseMock)

	usecaseMock.
		On("Request", mock.Anything).
		Return(errors.New("failed to request email sending"))

	handler := rest.NewSendEmailHandler(usecaseMock, logger)

	body := `{
		"to": "user@test.com",
		"subject": "Change password",
		"verification_code": "123456",
		"code_expiration_hours": "2"
	}`

	r := httptest.NewRequest(
		http.MethodPost,
		"/emails/change-password-code",
		strings.NewReader(body),
	)
	r.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	handler.SendChangePasswordCodeHandler(w, r)

	assertInternalServerError(t, w.Result(), usecaseMock)
}

// =============== Deletion code tests ===============
func TestSendDeletionCodeHandler_WhenRequestBodyIsInvalidJSON_ShouldReturnBadRequest(t *testing.T) {
	usecaseMock := new(RequestEmailUseCaseMock)

	handler := rest.NewSendEmailHandler(usecaseMock, logger)

	r := httptest.NewRequest(
		http.MethodPost,
		"/emails/deletion-code",
		strings.NewReader("{invalid-json"),
	)
	w := httptest.NewRecorder()

	handler.SendDeletionCodeHandler(w, r)

	assertBadRequest(t, w.Result(), usecaseMock)
}

func TestSendDeletionCodeHandler_WhenValidationFails_ShouldReturnUnprocessableEntity(t *testing.T) {
	usecaseMock := new(RequestEmailUseCaseMock)

	usecaseMock.
		On("Request", mock.Anything).
		Return(emailmessage.NewEmptyFieldError("verification_code"))

	handler := rest.NewSendEmailHandler(usecaseMock, logger)

	body := `{
		"to": "user@test.com",
		"subject": "Account deletion",
		"verification_code": "",
		"code_expiration_hours": "2"
	}`

	r := httptest.NewRequest(
		http.MethodPost,
		"/emails/deletion-code",
		strings.NewReader(body),
	)
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	handler.SendDeletionCodeHandler(w, r)

	assertUnprocessableEntity(t, w.Result(), usecaseMock)
}

func TestSendDeletionCodeHandler_WhenRequestIsValid_ShouldReturnAccepted(t *testing.T) {
	usecaseMock := new(RequestEmailUseCaseMock)

	usecaseMock.
		On("Request", mock.Anything).
		Return(nil)

	handler := rest.NewSendEmailHandler(usecaseMock, logger)

	body := `{
		"to": "user@test.com",
		"subject": "Account deletion",
		"verification_code": "123456",
		"code_expiration_hours": "2"
	}`

	r := httptest.NewRequest(
		http.MethodPost,
		"/emails/deletion-code",
		strings.NewReader(body),
	)
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	handler.SendDeletionCodeHandler(w, r)

	assertAccepted(t, w.Result(), usecaseMock)
}

func TestSendDeletionCodeHandler_WhenUnexpectedErrorOccurs_ShouldReturnInternalServerError(t *testing.T) {
	usecaseMock := new(RequestEmailUseCaseMock)

	usecaseMock.
		On("Request", mock.Anything).
		Return(errors.New("failed to request email sending"))

	handler := rest.NewSendEmailHandler(usecaseMock, logger)

	body := `{
		"to": "user@test.com",
		"subject": "Account deletion",
		"verification_code": "123456",
		"code_expiration_hours": "2"
	}`

	r := httptest.NewRequest(
		http.MethodPost,
		"/emails/deletion-code",
		strings.NewReader(body),
	)
	r.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	handler.SendDeletionCodeHandler(w, r)

	assertInternalServerError(t, w.Result(), usecaseMock)
}

// =============== Reset password code tests ===============
func TestSendResetPasswordCodeHandler_WhenRequestBodyIsInvalidJSON_ShouldReturnBadRequest(t *testing.T) {
	usecaseMock := new(RequestEmailUseCaseMock)

	handler := rest.NewSendEmailHandler(usecaseMock, logger)

	r := httptest.NewRequest(
		http.MethodPost,
		"/emails/reset-password-code",
		strings.NewReader("{invalid-json"),
	)
	w := httptest.NewRecorder()

	handler.SendResetPasswordCodeHandler(w, r)

	assertBadRequest(t, w.Result(), usecaseMock)
}

func TestSendResetPasswordCodeHandler_WhenValidationFails_ShouldReturnUnprocessableEntity(t *testing.T) {
	usecaseMock := new(RequestEmailUseCaseMock)

	usecaseMock.
		On("Request", mock.Anything).
		Return(emailmessage.NewEmptyFieldError("reset_password_link"))

	handler := rest.NewSendEmailHandler(usecaseMock, logger)

	body := `{
		"to": "user@test.com",
		"subject": "Reset password",
		"verification_code": "123456",
		"code_expiration_hours": "2",
		"reset_password_link": ""
	}`

	r := httptest.NewRequest(
		http.MethodPost,
		"/emails/reset-password-code",
		strings.NewReader(body),
	)
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	handler.SendResetPasswordCodeHandler(w, r)

	assertUnprocessableEntity(t, w.Result(), usecaseMock)
}

func TestSendResetPasswordCodeHandler_WhenRequestIsValid_ShouldReturnAccepted(t *testing.T) {
	usecaseMock := new(RequestEmailUseCaseMock)

	usecaseMock.
		On("Request", mock.Anything).
		Return(nil)

	handler := rest.NewSendEmailHandler(usecaseMock, logger)

	body := `{
		"to": "user@test.com",
		"subject": "Reset password",
		"verification_code": "123456",
		"code_expiration_hours": "2",
		"reset_password_link": "https://example.com/reset"
	}`

	r := httptest.NewRequest(
		http.MethodPost,
		"/emails/reset-password-code",
		strings.NewReader(body),
	)
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	handler.SendResetPasswordCodeHandler(w, r)

	assertAccepted(t, w.Result(), usecaseMock)
}

func TestSendResetPasswordCodeHandler_WhenUnexpectedErrorOccurs_ShouldReturnInternalServerError(t *testing.T) {
	usecaseMock := new(RequestEmailUseCaseMock)

	usecaseMock.
		On("Request", mock.Anything).
		Return(errors.New("failed to request email sending"))

	handler := rest.NewSendEmailHandler(usecaseMock, logger)

	body := `{
		"to": "user@test.com",
		"subject": "Reset password",
		"verification_code": "123456",
		"code_expiration_hours": "2",
		"reset_password_link": "https://example.com/reset"
	}`

	r := httptest.NewRequest(
		http.MethodPost,
		"/emails/reset-password-code",
		strings.NewReader(body),
	)
	r.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	handler.SendResetPasswordCodeHandler(w, r)

	assertInternalServerError(t, w.Result(), usecaseMock)
}
