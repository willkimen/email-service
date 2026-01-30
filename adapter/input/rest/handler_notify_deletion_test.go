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

func TestNotifyDeletionHandler_WhenRequestBodyIsInvalidJSON_ShouldReturnBadRequest(t *testing.T) {
	usecaseMock := new(RequestEmailUseCaseMock)

	handler := rest.NewSendEmailHandler(usecaseMock)

	r := httptest.NewRequest(
		http.MethodPost,
		"/emails/notify-deletion",
		strings.NewReader("{invalid-json"),
	)
	w := httptest.NewRecorder()

	handler.NotifyDeletionHandler(w, r)

	response := w.Result()

	require.Equal(t, http.StatusBadRequest, response.StatusCode)
	usecaseMock.AssertNotCalled(t, "Request", mock.Anything)
}

func TestNotifyDeletionHandler_WhenValidationFails_ShouldReturnUnprocessableEntity(t *testing.T) {
	usecaseMock := new(RequestEmailUseCaseMock)

	usecaseMock.
		On("Request", mock.Anything).
		Return(emailmessage.NewEmptyFieldError("to"))

	handler := rest.NewSendEmailHandler(usecaseMock)

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

	require.Equal(t, http.StatusUnprocessableEntity, response.StatusCode)
	usecaseMock.AssertCalled(t, "Request", mock.Anything)
}

func TestNotifyDeletionHandler_WhenRequestIsValid_ShouldReturnAccepted(t *testing.T) {
	usecaseMock := new(RequestEmailUseCaseMock)

	usecaseMock.
		On("Request", mock.Anything).
		Return(nil)

	handler := rest.NewSendEmailHandler(usecaseMock)

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

	require.Equal(t, http.StatusAccepted, response.StatusCode)
	usecaseMock.AssertCalled(t, "Request", mock.Anything)
}

func TestNotifyDeletionHandler_WhenUnexpectedErrorOccurs_ShouldReturnInternalServerError(t *testing.T) {
	usecaseMock := new(RequestEmailUseCaseMock)

	usecaseMock.
		On("Request", mock.Anything).
		Return(errors.New("failed to request email sending"))

	handler := rest.NewSendEmailHandler(usecaseMock)

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

	require.Equal(t, http.StatusInternalServerError, response.StatusCode)
	usecaseMock.AssertCalled(t, "Request", mock.Anything)
}
