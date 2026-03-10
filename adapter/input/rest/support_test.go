package rest_test

import (
	"emailservice/core/application/email_message"
	"encoding/json"
	"log/slog"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

var logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))

type RequestEmailUseCaseMock struct {
	mock.Mock
}

func (m *RequestEmailUseCaseMock) Request(message emailmessage.EmailMessage) error {
	args := m.Called(message)
	return args.Error(0)
}

func decodeJSONResponse(t *testing.T, res *http.Response, target any) {
	t.Helper()

	defer res.Body.Close()

	err := json.NewDecoder(res.Body).Decode(target)
	require.NoError(t, err)
}

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

func assertFieldValidationError(
	t *testing.T,
	response map[string]string,
	expectedError string,
	expectedField string,
) {
	t.Helper()

	assert.Equal(t, expectedError, response["error"])
	assert.Equal(t, expectedField, response["field"])
}
