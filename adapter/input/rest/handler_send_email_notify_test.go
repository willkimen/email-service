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

	assertBadRequest(t, w.Result(), usecaseMock)
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

	assertUnprocessableEntity(t, w.Result(), usecaseMock)
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

	assertAccepted(t, w.Result(), usecaseMock)
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

	assertInternalServerError(t, w.Result(), usecaseMock)
}

func TestNotifyActivationHandler_WhenEmptyField_ShouldReturnValidationError(t *testing.T) {
	tests := []struct {
		name          string
		field         string
		expectedError string
	}{
		{
			name:          "missing to",
			field:         "to",
			expectedError: "to field is required",
		},
		{
			name:          "missing subject",
			field:         "subject",
			expectedError: "subject field is required",
		},
		{
			name:          "missing login_link",
			field:         "login_link",
			expectedError: "login_link field is required",
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {

			usecaseMock := new(RequestEmailUseCaseMock)

			usecaseMock.
				On("Request", mock.Anything).
				Return(emailmessage.NewEmptyFieldError(tt.field))

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

			res := w.Result()

			assert.Equal(t, http.StatusUnprocessableEntity, res.StatusCode)

			var response map[string]string
			decodeJSONResponse(t, res, &response)

			assertFieldValidationError(
				t,
				response,
				tt.expectedError,
				tt.field,
			)

			usecaseMock.AssertCalled(t, "Request", mock.Anything)
		})
	}
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

	assertBadRequest(t, w.Result(), usecaseMock)
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

	assertUnprocessableEntity(t, w.Result(), usecaseMock)
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

	assertAccepted(t, w.Result(), usecaseMock)
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

	assertInternalServerError(t, w.Result(), usecaseMock)
}

func TestNotifyChangeEmailHandler_WhenEmptyField_ShouldReturnValidationError(t *testing.T) {
	tests := []struct {
		name          string
		field         string
		expectedError string
	}{
		{
			name:          "missing to",
			field:         "to",
			expectedError: "to field is required",
		},
		{
			name:          "missing subject",
			field:         "subject",
			expectedError: "subject field is required",
		},
		{
			name:          "missing login_link",
			field:         "login_link",
			expectedError: "login_link field is required",
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {

			usecaseMock := new(RequestEmailUseCaseMock)

			usecaseMock.
				On("Request", mock.Anything).
				Return(emailmessage.NewEmptyFieldError(tt.field))

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

			res := w.Result()

			assert.Equal(t, http.StatusUnprocessableEntity, res.StatusCode)

			var response map[string]string
			decodeJSONResponse(t, res, &response)

			assertFieldValidationError(
				t,
				response,
				tt.expectedError,
				tt.field,
			)

			usecaseMock.AssertCalled(t, "Request", mock.Anything)
		})
	}
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

	assertBadRequest(t, w.Result(), usecaseMock)
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

	assertUnprocessableEntity(t, w.Result(), usecaseMock)
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

	assertAccepted(t, w.Result(), usecaseMock)
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

	assertInternalServerError(t, w.Result(), usecaseMock)
}

func TestNotifyChangePasswordHandler_WhenEmptyField_ShouldReturnValidationError(t *testing.T) {
	tests := []struct {
		name          string
		field         string
		expectedError string
	}{
		{
			name:          "missing to",
			field:         "to",
			expectedError: "to field is required",
		},
		{
			name:          "missing subject",
			field:         "subject",
			expectedError: "subject field is required",
		},
		{
			name:          "missing login_link",
			field:         "login_link",
			expectedError: "login_link field is required",
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {

			usecaseMock := new(RequestEmailUseCaseMock)

			usecaseMock.
				On("Request", mock.Anything).
				Return(emailmessage.NewEmptyFieldError(tt.field))

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

			res := w.Result()

			assert.Equal(t, http.StatusUnprocessableEntity, res.StatusCode)

			var response map[string]string
			decodeJSONResponse(t, res, &response)

			assertFieldValidationError(
				t,
				response,
				tt.expectedError,
				tt.field,
			)

			usecaseMock.AssertCalled(t, "Request", mock.Anything)
		})
	}
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

	assertBadRequest(t, w.Result(), usecaseMock)
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

	assertUnprocessableEntity(t, w.Result(), usecaseMock)
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

	assertAccepted(t, w.Result(), usecaseMock)
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

	assertInternalServerError(t, w.Result(), usecaseMock)
}

func TestNotifyDeletionHandler_WhenEmptyField_ShouldReturnValidationError(t *testing.T) {
	tests := []struct {
		name          string
		field         string
		expectedError string
	}{
		{
			name:          "missing to",
			field:         "to",
			expectedError: "to field is required",
		},
		{
			name:          "missing subject",
			field:         "subject",
			expectedError: "subject field is required",
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {

			usecaseMock := new(RequestEmailUseCaseMock)

			usecaseMock.
				On("Request", mock.Anything).
				Return(emailmessage.NewEmptyFieldError(tt.field))

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

			res := w.Result()

			assert.Equal(t, http.StatusUnprocessableEntity, res.StatusCode)

			var response map[string]string
			decodeJSONResponse(t, res, &response)

			assertFieldValidationError(
				t,
				response,
				tt.expectedError,
				tt.field,
			)

			usecaseMock.AssertCalled(t, "Request", mock.Anything)
		})
	}
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

	assertBadRequest(t, w.Result(), usecaseMock)
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

	assertUnprocessableEntity(t, w.Result(), usecaseMock)
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

	assertAccepted(t, w.Result(), usecaseMock)
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

	assertInternalServerError(t, w.Result(), usecaseMock)
}

func TestNotifyResetPasswordHandler_WhenEmptyField_ShouldReturnValidationError(t *testing.T) {
	tests := []struct {
		name          string
		field         string
		expectedError string
	}{
		{
			name:          "missing to",
			field:         "to",
			expectedError: "to field is required",
		},
		{
			name:          "missing subject",
			field:         "subject",
			expectedError: "subject field is required",
		},
		{
			name:          "missing login_link",
			field:         "login_link",
			expectedError: "login_link field is required",
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {

			usecaseMock := new(RequestEmailUseCaseMock)

			usecaseMock.
				On("Request", mock.Anything).
				Return(emailmessage.NewEmptyFieldError(tt.field))

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

			res := w.Result()

			assert.Equal(t, http.StatusUnprocessableEntity, res.StatusCode)

			var response map[string]string
			decodeJSONResponse(t, res, &response)

			assertFieldValidationError(
				t,
				response,
				tt.expectedError,
				tt.field,
			)

			usecaseMock.AssertCalled(t, "Request", mock.Anything)
		})
	}
}
