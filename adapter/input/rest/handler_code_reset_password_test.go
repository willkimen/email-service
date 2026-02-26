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

	response := w.Result()

	assert.Equal(t, http.StatusBadRequest, response.StatusCode,
		"expected status code to be 400 when request body contains invalid JSON")

	usecaseMock.AssertNotCalled(t, "Request", mock.Anything)
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

	response := w.Result()

	assert.Equal(t, http.StatusUnprocessableEntity, response.StatusCode,
		"expected status code to be 422 when validation fails")

	usecaseMock.AssertCalled(t, "Request", mock.Anything)
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

	response := w.Result()

	assert.Equal(t, http.StatusAccepted, response.StatusCode,
		"expected status code to be 202 when request is valid")

	usecaseMock.AssertCalled(t, "Request", mock.Anything)
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

	response := w.Result()

	assert.Equal(t, http.StatusInternalServerError, response.StatusCode,
		"expected status code to be 500 when unexpected error occurs")

	usecaseMock.AssertCalled(t, "Request", mock.Anything)
}
