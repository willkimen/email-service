package rest_test

import (
	"emailservice/adapter/input/rest"
	"emailservice/core/application/apperrors"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSendEmailHandler_ReturnAccepted(t *testing.T) {
	// arrange
	usecaseMock := new(RequestSendEmailMock)
	usecaseMock.On("Execute", mock.Anything).Return(nil)

	handler := rest.NewSendEmailHandler(usecaseMock, loggerDummy)

	body := map[string]any{
		"id":   "fake-id",
		"to":   "email@email.com",
		"type": "email_verification_code",
		"variables": map[string]string{
			"verification_code": "123456",
		},
	}
	bodyBytes, err := json.Marshal(body)
	assert.NoError(t, err)

	r := httptest.NewRequest(
		http.MethodPost, "/api/v1/email", strings.NewReader(string(bodyBytes)),
	)
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// act
	handler.SendEmailHandler(w, r)

	// arrange
	res := w.Result()

	assert.Equal(t, http.StatusAccepted, res.StatusCode)
	assert.Contains(t, w.Body.String(), "accepted")

	usecaseMock.AssertCalled(t, "Execute", mock.Anything)
}

func TestSendEmailHandler_EmailFormatIsInvalid_ReturnError(t *testing.T) {
	// Arrange
	expected_err := apperrors.NewEmailInvalidFormatError("to")
	usecaseMock := new(RequestSendEmailMock)
	handler := rest.NewSendEmailHandler(usecaseMock, loggerDummy)
	bodyWithEmailFormatInvalid := map[string]any{
		"id":   "fake-id",
		"to":   "invalid-email",
		"type": "email_verification_code",
		"variables": map[string]string{
			"verification_code": "123456",
		},
	}
	bodyBytes, err := json.Marshal(bodyWithEmailFormatInvalid)
	assert.NoError(t, err)

	r := httptest.NewRequest(
		http.MethodPost, "/api/v1/email", strings.NewReader(string(bodyBytes)),
	)
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// Act
	handler.SendEmailHandler(w, r)

	// Asserts
	res := w.Result()
	var response map[string]any
	decodeJSONResponse(t, res, &response)

	assert.Equal(t, http.StatusUnprocessableEntity, res.StatusCode)
	assertInvalidFieldError(t, response, expected_err)
	usecaseMock.AssertNotCalled(t, "Execute", mock.Anything)
}

func TestSendEmailHandler_VerificationCodeKeyMissing_ReturnError(t *testing.T) {
	// Arrange
	expected_err := apperrors.NewVerificationCodeMissingError(
		"variables",
		"variables cannot be null or empty for this message type",
	)
	usecaseMock := new(RequestSendEmailMock)

	handler := rest.NewSendEmailHandler(usecaseMock, loggerDummy)

	bodyWithoutVariables := map[string]any{
		"id":   "fake-id",
		"to":   "email@email.com",
		"type": "email_verification_code",
	}
	bodyBytes, err := json.Marshal(bodyWithoutVariables)
	assert.NoError(t, err)

	r := httptest.NewRequest(
		http.MethodPost, "/api/v1/email", strings.NewReader(string(bodyBytes)),
	)
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// Act
	handler.SendEmailHandler(w, r)

	// Asserts
	res := w.Result()
	var response map[string]any
	decodeJSONResponse(t, res, &response)

	assert.Equal(t, http.StatusUnprocessableEntity, res.StatusCode)
	assertInvalidFieldError(t, response, expected_err)
	usecaseMock.AssertNotCalled(t, "Execute", mock.Anything)
}
func TestSendEmailHandler_TypeInvalid_ReturnError(t *testing.T) {
	// Arrange
	expected_err := apperrors.NewInvalidTypeError("type")
	usecaseMock := new(RequestSendEmailMock)
	handler := rest.NewSendEmailHandler(usecaseMock, loggerDummy)

	bodyWithInvalidType := map[string]any{
		"id":   "fake-id",
		"to":   "email@email.com",
		"type": "type-invalid",
		"variables": map[string]string{
			"verification_code": "123456",
		},
	}
	bodyBytes, err := json.Marshal(bodyWithInvalidType)
	assert.NoError(t, err)

	r := httptest.NewRequest(
		http.MethodPost, "/api/v1/email", strings.NewReader(string(bodyBytes)),
	)
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// Act
	handler.SendEmailHandler(w, r)

	// Asserts
	res := w.Result()
	var response map[string]any
	decodeJSONResponse(t, res, &response)

	assert.Equal(t, http.StatusUnprocessableEntity, res.StatusCode)
	assertInvalidFieldError(t, response, expected_err)
}

func TestSendEmailHandler_EmptyField_ReturnError(t *testing.T) {
	for _, c := range emptyFieldCases {
		t.Run(c.name, func(t *testing.T) {
			// arrange
			usecaseMock := new(RequestSendEmailMock)
			handler := rest.NewSendEmailHandler(usecaseMock, loggerDummy)

			bodyBytes, err := json.Marshal(c.bodyWithEmptyField)
			assert.NoError(t, err)

			r := httptest.NewRequest(
				http.MethodPost,
				"/api/v1/email",
				strings.NewReader(string(bodyBytes)),
			)
			r.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			// act
			handler.SendEmailHandler(w, r)

			// assert
			res := w.Result()
			var response map[string]any
			decodeJSONResponse(t, res, &response)

			assert.Equal(t, http.StatusUnprocessableEntity, res.StatusCode)
			assertInvalidFieldError(t, response, c.expectedError)
			usecaseMock.AssertNotCalled(t, "Execute", mock.Anything)
		})
	}
}

func TestSendEmailHandler_InvalidJSON_ReturnBadRequest(t *testing.T) {
	// arrgange
	usecaseMock := new(RequestSendEmailMock)
	handler := rest.NewSendEmailHandler(usecaseMock, loggerDummy)

	r := httptest.NewRequest(
		http.MethodPost, "/api/v1/email", strings.NewReader("{invalid-json"),
	)
	w := httptest.NewRecorder()

	// act
	handler.SendEmailHandler(w, r)

	// assert
	res := w.Result()
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	assert.Contains(t, w.Body.String(), "error")

	usecaseMock.AssertNotCalled(t, "Execute", mock.Anything)
}

func TestSendEmailHandler_UnexpectedError_ReturnInternalServerError(t *testing.T) {
	// arrange
	expectedError := apperrors.NewInfrastructureError(
		"Un unexpected Error ocurred",
		errors.New("failed to request email sending"),
	)
	usecaseMock := new(RequestSendEmailMock)
	usecaseMock.On("Execute", mock.Anything).Return(expectedError)
	handler := rest.NewSendEmailHandler(usecaseMock, loggerDummy)

	body := map[string]any{
		"id":   "fake-id",
		"to":   "email@email.com",
		"type": "email_verification_code",
		"variables": map[string]string{
			"verification_code": "123456",
		},
	}
	bodyBytes, err := json.Marshal(body)
	assert.NoError(t, err)

	r := httptest.NewRequest(
		http.MethodPost, "/api/v1/email", strings.NewReader(string(bodyBytes)),
	)
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// act
	handler.SendEmailHandler(w, r)

	res := w.Result()
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
	assert.Contains(t, w.Body.String(), "internal server error")

	usecaseMock.AssertCalled(t, "Execute", mock.Anything)
}
