package rest

import (
	"emailservice/core/application/email_message"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// envelope represents a generic JSON response structure.
type envelope map[string]any

// writeJSON serializes the given envelope into JSON.
func (s *SendEmailHandler) writeJSON(data envelope) ([]byte, error) {
	js, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	return js, nil
}

// writeResponse writes the HTTP status code and JSON body to the response.
func (s *SendEmailHandler) writeResponse(w http.ResponseWriter, status int, json []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(json)
}

// serverErrorResponse sends a generic 500 response.
//
// This is used as a defensive fallback for unexpected infrastructure failures.
func (s *SendEmailHandler) serverErrorResponse(w http.ResponseWriter, r *http.Request) {
	http.Error(
		w,
		http.StatusText(http.StatusInternalServerError),
		http.StatusInternalServerError,
	)
}

// respond writes a JSON response to the client.
//
// If JSON serialization fails, a 500 Internal Server Error is returned.
// This situation represents an unexpected infrastructure failure.
func (s *SendEmailHandler) respond(
	w http.ResponseWriter,
	r *http.Request,
	status int,
	data envelope,
) {
	json, err := s.writeJSON(data)
	if err != nil {
		s.serverErrorResponse(w, r)
		return
	}
	s.writeResponse(w, status, json)
}

// readJSON decodes and validates the JSON request body.
//
// It enforces size limits, disallows unknown fields, and ensures
// that the request body contains exactly one JSON value.
func (s *SendEmailHandler) readJSON(w http.ResponseWriter, r *http.Request, dto any) error {
	maxBytes := 1_048_576
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	err := dec.Decode(dto)
	if err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError
		var invalidUnmarshalError *json.InvalidUnmarshalError
		var maxBytesError *http.MaxBytesError

		switch {
		case errors.As(err, &syntaxError):
			return fmt.Errorf("body contains badly-formed JSON (at character %d)", syntaxError.Offset)

		case errors.Is(err, io.ErrUnexpectedEOF):
			return errors.New("body contains badly-formed JSON")

		case errors.As(err, &unmarshalTypeError):
			if unmarshalTypeError.Field != "" {
				return fmt.Errorf("body contains an incorrect JSON type for field %q", unmarshalTypeError.Field)
			}
			return fmt.Errorf("body contains an incorrect JSON type (at character %d)", unmarshalTypeError.Offset)

		case errors.Is(err, io.EOF):
			return errors.New("body must not be empty")

		case strings.HasPrefix(err.Error(), "json: unknown field "):
			fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
			return fmt.Errorf("body contains unknown key %s", fieldName)

		case errors.As(err, &maxBytesError):
			return fmt.Errorf("body must not be larger than %d bytes", maxBytesError.Limit)

		case errors.As(err, &invalidUnmarshalError):
			panic(err)

		default:
			return err
		}
	}

	err = dec.Decode(&struct{}{})
	if !errors.Is(err, io.EOF) {
		return errors.New("body must contain only a single JSON value")
	}

	return nil
}

// handleEmailRequest is a generic HTTP handler for email-related requests.
//
// It is responsible for:
// - Reading and validating the JSON payload
// - Converting the DTO into an email message
// - Delegating the request to the application use case
// - Mapping infrastructure errors to HTTP responses
func (s *SendEmailHandler) handleEmailRequest(w http.ResponseWriter, r *http.Request, dto EmailRequestDTO) {
	// Possible responses:
	// - 202 Accepted: request accepted for asynchronous processing
	// - 400 Bad Request: malformed or invalid JSON payload
	// - 422 Unprocessable Entity: field validation error
	// - 500 Internal Server Error: unexpected internal failure
	s.Logger.InfoContext(
		r.Context(),
		"http email request received",
		"method", r.Method,
		"path", r.URL.Path,
	)

	err := s.readJSON(w, r, dto)
	if err != nil {
		s.Logger.ErrorContext(
			r.Context(),
			"invalid json payload",
			"error", err,
			"method", r.Method,
			"path", r.URL.Path,
		)

		s.respond(w, r, http.StatusBadRequest, envelope{"error": err.Error()})
		return
	}

	err = s.Usecase.Request(dto.ToEmailMessage())
	if err != nil {
		var validationErr emailmessage.FieldValidationError

		if errors.As(err, &validationErr) {
			s.Logger.ErrorContext(
				r.Context(),
				"email validation error",
				"error", err,
				"method", r.Method,
				"path", r.URL.Path,
			)

			s.respond(
				w,
				r,
				http.StatusUnprocessableEntity,
				envelope{
					"error": validationErr.Error(),
					"field": validationErr.GetField(),
				},
			)
			return
		}

		s.Logger.ErrorContext(
			r.Context(),
			"internal error while requesting email",
			"error", err,
			"method", r.Method,
			"path", r.URL.Path,
		)

		s.respond(
			w,
			r,
			http.StatusInternalServerError,
			envelope{"error": "internal server error"},
		)
		return
	}

	s.Logger.InfoContext(
		r.Context(),
		"email request accepted",
		"method", r.Method,
		"path", r.URL.Path,
	)

	s.respond(w, r, http.StatusAccepted, envelope{"status": "accepted"})
}
