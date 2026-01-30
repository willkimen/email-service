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

	require.Error(t, err)
	assert.Contains(t, err.Error(), "badly-formed JSON")
}

func TestReadJSON_WhenJSONHasUnexpectedEOF_ShouldReturnMalformedJSONError(t *testing.T) {
	handler := SendEmailHandler{}

	var dto testDTO
	req := newJSONRequest(`{"name": "john"`)
	w := httptest.NewRecorder()

	err := handler.readJSON(w, req, &dto)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "badly-formed JSON")
}

func TestReadJSON_WhenFieldHasWrongType_ShouldReturnTypeError(t *testing.T) {
	handler := SendEmailHandler{}

	var dto testDTO
	req := newJSONRequest(`{"name": "john", "age": "wrong"}`)
	w := httptest.NewRecorder()

	err := handler.readJSON(w, req, &dto)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "incorrect JSON type")
}

func TestReadJSON_WhenBodyIsEmpty_ShouldReturnEmptyBodyError(t *testing.T) {
	handler := SendEmailHandler{}

	var dto testDTO
	req := newJSONRequest(``)
	w := httptest.NewRecorder()

	err := handler.readJSON(w, req, &dto)

	require.Error(t, err)
	assert.Equal(t, "body must not be empty", err.Error())
}

func TestReadJSON_WhenUnknownFieldIsPresent_ShouldReturnUnknownFieldError(t *testing.T) {
	handler := SendEmailHandler{}

	var dto testDTO
	req := newJSONRequest(`{"name": "john", "unknown": "x"}`)
	w := httptest.NewRecorder()

	err := handler.readJSON(w, req, &dto)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "body contains unknown key")
}

func TestReadJSON_WhenBodyExceedsMaxSize_ShouldReturnSizeError(t *testing.T) {
	handler := SendEmailHandler{}

	var dto testDTO
	largeBody := strings.Repeat("a", 1_048_577)
	req := newJSONRequest(`{"name":"` + largeBody + `"}`)
	w := httptest.NewRecorder()

	err := handler.readJSON(w, req, &dto)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "must not be larger than")
}

func TestReadJSON_WhenMultipleJSONValuesProvided_ShouldReturnSingleValueError(t *testing.T) {
	handler := SendEmailHandler{}

	var dto testDTO
	req := newJSONRequest(`{"name":"john","age":30} {"name":"doe","age":40}`)
	w := httptest.NewRecorder()

	err := handler.readJSON(w, req, &dto)

	require.Error(t, err)
	assert.Equal(t, "body must contain only a single JSON value", err.Error())
}

func TestReadJSON_WhenJSONIsValid_ShouldDecodeSuccessfully(t *testing.T) {
	handler := SendEmailHandler{}

	var dto testDTO
	req := newJSONRequest(`{"name":"john","age":30}`)
	w := httptest.NewRecorder()

	err := handler.readJSON(w, req, &dto)

	require.NoError(t, err)
	assert.Equal(t, "john", dto.Name)
	assert.Equal(t, 30, dto.Age)
}
