package rest_test

import (
	"emailservice/core/application/apperrors"
	"emailservice/core/application/message"
	"encoding/json"
	"log/slog"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

var loggerDummy = slog.New(slog.NewJSONHandler(os.Stdout, nil))

type RequestSendEmailMock struct {
	mock.Mock
}

func (m *RequestSendEmailMock) Execute(message message.Message) error {
	args := m.Called(message)
	return args.Error(0)
}

var emptyFieldCases = []struct {
	name               string
	bodyWithEmptyField map[string]any
	expectedError      error
}{
	{
		name: "missing id",
		bodyWithEmptyField: map[string]any{
			"to":   "email@email.com",
			"type": "email_verification_code",
			"variables": map[string]string{
				"verification_code": "123456",
			},
		},
		expectedError: apperrors.NewEmptyFieldError("id"),
	},
	{
		name: "missing to",
		bodyWithEmptyField: map[string]any{
			"id":   "fake-id",
			"type": "email_verification_code",
			"variables": map[string]string{
				"verification_code": "123456",
			},
		},
		expectedError: apperrors.NewEmptyFieldError("to"),
	},
	{
		name: "missing type",
		bodyWithEmptyField: map[string]any{
			"id": "fake-id",
			"to": "email@email.com",
			"variables": map[string]string{
				"verification_code": "123456",
			},
		},
		expectedError: apperrors.NewEmptyFieldError("type"),
	},
}

// ========================  auxiliary functions for assertions ==========================
func assertInvalidFieldError(t *testing.T, response map[string]any, err error) {
	t.Helper()

	var invalidField *apperrors.InvalidFieldError
	if assert.ErrorAs(t, err, &invalidField) {
		assert.Equal(t, invalidField.Error(), response["error"])
		assert.Equal(t, invalidField.FieldName, response["field"])
	}
}

func decodeJSONResponse(t *testing.T, res *http.Response, target any) {
	t.Helper()

	defer res.Body.Close()

	err := json.NewDecoder(res.Body).Decode(target)
	require.NoError(t, err)
}
