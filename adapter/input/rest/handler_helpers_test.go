package rest

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type testDTO struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func newJSONRequest(body string) *http.Request {
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	return req
}

func TestReadJSON_WhenJSONIsMalformed_ShouldReturnSyntaxError(t *testing.T) {
	handler := SendEmailHandler{}

	var dto testDTO
	req := newJSONRequest(`{"name": "john",`)
	w := httptest.NewRecorder()

	err := handler.readJSON(w, req, &dto)

	require.Error(t, err,
		"expected readJSON to return error when JSON is malformed")

	assert.Contains(t, err.Error(), "badly-formed JSON",
		"expected error message to indicate badly-formed JSON")
}

func TestReadJSON_WhenJSONHasUnexpectedEOF_ShouldReturnMalformedJSONError(t *testing.T) {
	handler := SendEmailHandler{}

	var dto testDTO
	req := newJSONRequest(`{"name": "john"`)
	w := httptest.NewRecorder()

	err := handler.readJSON(w, req, &dto)

	require.Error(t, err,
		"expected readJSON to return error when JSON has unexpected EOF")

	assert.Contains(t, err.Error(), "badly-formed JSON",
		"expected error message to indicate badly-formed JSON")
}

func TestReadJSON_WhenFieldHasWrongType_ShouldReturnTypeError(t *testing.T) {
	handler := SendEmailHandler{}

	var dto testDTO
	req := newJSONRequest(`{"name": "john", "age": "wrong"}`)
	w := httptest.NewRecorder()

	err := handler.readJSON(w, req, &dto)

	require.Error(t, err,
		"expected readJSON to return error when a field has wrong type")

	assert.Contains(t, err.Error(), "incorrect JSON type",
		"expected error message to indicate incorrect JSON type")
}

func TestReadJSON_WhenBodyIsEmpty_ShouldReturnEmptyBodyError(t *testing.T) {
	handler := SendEmailHandler{}

	var dto testDTO
	req := newJSONRequest(``)
	w := httptest.NewRecorder()

	err := handler.readJSON(w, req, &dto)

	require.Error(t, err,
		"expected readJSON to return error when body is empty")

	assert.Equal(t, "body must not be empty", err.Error(),
		"expected error message to indicate empty body")
}

func TestReadJSON_WhenUnknownFieldIsPresent_ShouldReturnUnknownFieldError(t *testing.T) {
	handler := SendEmailHandler{}

	var dto testDTO
	req := newJSONRequest(`{"name": "john", "unknown": "x"}`)
	w := httptest.NewRecorder()

	err := handler.readJSON(w, req, &dto)

	require.Error(t, err,
		"expected readJSON to return error when unknown field is present")

	assert.Contains(t, err.Error(), "body contains unknown key",
		"expected error message to indicate unknown key")
}

func TestReadJSON_WhenBodyExceedsMaxSize_ShouldReturnSizeError(t *testing.T) {
	handler := SendEmailHandler{}

	var dto testDTO
	largeBody := strings.Repeat("a", 1_048_577)
	req := newJSONRequest(`{"name":"` + largeBody + `"}`)
	w := httptest.NewRecorder()

	err := handler.readJSON(w, req, &dto)

	require.Error(t, err,
		"expected readJSON to return error when body exceeds maximum size")

	assert.Contains(t, err.Error(), "must not be larger than",
		"expected error message to indicate body size limit")
}

func TestReadJSON_WhenMultipleJSONValuesProvided_ShouldReturnSingleValueError(t *testing.T) {
	handler := SendEmailHandler{}

	var dto testDTO
	req := newJSONRequest(`{"name":"john","age":30} {"name":"doe","age":40}`)
	w := httptest.NewRecorder()

	err := handler.readJSON(w, req, &dto)

	require.Error(t, err,
		"expected readJSON to return error when multiple JSON values are provided")

	assert.Equal(t, "body must contain only a single JSON value", err.Error(),
		"expected error message to indicate single JSON value requirement")
}

func TestReadJSON_WhenJSONIsValid_ShouldDecodeSuccessfully(t *testing.T) {
	handler := SendEmailHandler{}

	var dto testDTO
	req := newJSONRequest(`{"name":"john","age":30}`)
	w := httptest.NewRecorder()

	err := handler.readJSON(w, req, &dto)

	require.NoError(t, err,
		"expected readJSON to return nil when JSON is valid")

	assert.Equal(t, "john", dto.Name,
		"expected Name field to be decoded correctly")

	assert.Equal(t, 30, dto.Age,
		"expected Age field to be decoded correctly")
}
