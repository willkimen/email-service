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

type envelope map[string]any

func (h *HandlerEmail) writeJSON(data envelope) ([]byte, error) {
	js, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	return js, nil
}

func (h *HandlerEmail) writeResponse(w http.ResponseWriter, status int, json []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(json)
}

func (h *HandlerEmail) serverErrorResponse(w http.ResponseWriter, r *http.Request) {
	http.Error(
		w,
		http.StatusText(http.StatusInternalServerError),
		http.StatusInternalServerError,
	)
}

// respond writes a JSON response to the client.
//
// If JSON serialization fails, a 500 Internal Server Error
// is returned as a defensive fallback. This scenario represents
// an unexpected infrastructure failure, not a business error.
func (h *HandlerEmail) respond(
	w http.ResponseWriter,
	r *http.Request,
	status int,
	data envelope,
) {
	json, err := h.writeJSON(data)
	if err != nil {
		h.serverErrorResponse(w, r)
		return
	}
	h.writeResponse(w, status, json)
}

func (h *HandlerEmail) readJSON(w http.ResponseWriter, r *http.Request, dto any) error {
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

func (h *HandlerEmail) handleEmailRequest(w http.ResponseWriter, r *http.Request, dto EmailRequestDTO) {
	// Possible responses:
	// - 202 Accepted: request was accepted for asynchronous processing
	// - 400 Bad Request: malformed or invalid JSON payload
	// - 422 Unprocessable Entity: domain validation error
	// - 500 Internal Server Error: fallback response for unexpected internal failures
	//
	// Note:
	// The handler does not actively decide when to return 500.
	// InternalServerError is used only as a last-resort fallback
	// when response serialization or infrastructure unexpectedly fails.
	err := h.readJSON(w, r, dto)
	if err != nil {
		h.respond(w, r, http.StatusBadRequest, envelope{"error": err.Error()})
		return
	}

	err = h.Usecase.Request(dto.ToEmailMessage())
	if err != nil {
		var validationErr emailmessage.FieldValidationError

		if errors.As(err, &validationErr) {
			h.respond(
				w,
				r,
				http.StatusUnprocessableEntity,
				envelope{"error": err.Error()},
			)
			return
		}

		h.respond(
			w,
			r,
			http.StatusInternalServerError,
			envelope{"error": "internal server error"},
		)
		return
	}
	h.respond(w, r, http.StatusAccepted, envelope{"status": "accepted"})
}
