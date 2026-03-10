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

// ======================= Test invalid email format ==============================
func TestWhenEmailFormatIsInvalid_ShouldReturnValidationError(t *testing.T) {
	usecaseMock := new(RequestEmailUseCaseMock)

	usecaseMock.
		On("Request", mock.Anything).
		Return(emailmessage.NewEmailInvalidFormatError())

	handler := rest.NewSendEmailHandler(usecaseMock, logger)

	body := `{
		"to": "invalid-email",
		"subject": "Activate your account",
		"verification_code": "123456",
		"code_expiration_hours": "2",
		"activation_link": "https://example.com/activate",
		"activation_deadline_days": "7"
	}`

	r := httptest.NewRequest(
		http.MethodPost,
		"/emails/activation-code",
		strings.NewReader(body),
	)

	r.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	handler.SendActivationCodeHandler(w, r)

	res := w.Result()

	assert.Equal(t, http.StatusUnprocessableEntity, res.StatusCode)

	var response map[string]string
	decodeJSONResponse(t, res, &response)

	assertFieldValidationError(
		t,
		response,
		"email format is invalid",
		"",
	)
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

func TestSendActivationCodeHandler_WhenEmptyField_ShouldReturnValidationError(t *testing.T) {
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
			name:          "missing verification_code",
			field:         "verification_code",
			expectedError: "verification_code field is required",
		},
		{
			name:          "missing code_expiration_hours",
			field:         "code_expiration_hours",
			expectedError: "code_expiration_hours field is required",
		},
		{
			name:          "missing activation_link",
			field:         "activation_link",
			expectedError: "activation_link field is required",
		},
		{
			name:          "missing activation_deadline_days",
			field:         "activation_deadline_days",
			expectedError: "activation_deadline_days field is required",
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
				"subject": "Activate your account",
				"verification_code": "123456",
				"code_expiration_hours": "2",
				"activation_link": "https://example.com/activate",
				"activation_deadline_days": "7"
			}`

			r := httptest.NewRequest(
				http.MethodPost,
				"/emails/activation-code",
				strings.NewReader(body),
			)

			r.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()

			handler.SendActivationCodeHandler(w, r)

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
		})
	}
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

func TestSendChangeEmailCodeHandler_WhenEmptyField_ShouldReturnValidationError(t *testing.T) {
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
			name:          "missing verification_code",
			field:         "verification_code",
			expectedError: "verification_code field is required",
		},
		{
			name:          "missing code_expiration_hours",
			field:         "code_expiration_hours",
			expectedError: "code_expiration_hours field is required",
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
		})
	}
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

func TestSendChangePasswordCodeHandler_WhenEmptyField_ShouldReturnValidationError(t *testing.T) {
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
			name:          "missing verification_code",
			field:         "verification_code",
			expectedError: "verification_code field is required",
		},
		{
			name:          "missing code_expiration_hours",
			field:         "code_expiration_hours",
			expectedError: "code_expiration_hours field is required",
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
		})
	}
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

func TestSendDeletionCodeHandler_WhenEmptyField_ShouldReturnValidationError(t *testing.T) {
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
			name:          "missing verification_code",
			field:         "verification_code",
			expectedError: "verification_code field is required",
		},
		{
			name:          "missing code_expiration_hours",
			field:         "code_expiration_hours",
			expectedError: "code_expiration_hours field is required",
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
		})
	}
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

func TestSendResetPasswordCodeHandler_WhenEmptyField_ShouldReturnValidationError(t *testing.T) {
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
			name:          "missing verification_code",
			field:         "verification_code",
			expectedError: "verification_code field is required",
		},
		{
			name:          "missing code_expiration_hours",
			field:         "code_expiration_hours",
			expectedError: "code_expiration_hours field is required",
		},
		{
			name:          "missing reset_password_link",
			field:         "reset_password_link",
			expectedError: "reset_password_link field is required",
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
		})
	}
}
