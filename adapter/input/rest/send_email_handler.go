package rest

import (
	"emailservice/core/application/apperrors"
	"emailservice/core/application/message"
	"emailservice/core/application/ports/input"
	"errors"
	"log/slog"
	"net/http"
)

// SendEmailHandler handles HTTP requests related to email sending operations.
//
// It acts as an input adapter, receiving HTTP requests, converting them
// into messages, and delegating execution to the appropriate
// application use case.
type SendEmailHandler struct {
	RequestSendEmailUsecase inputport.RequestSendEmailPort
	Logger                  *slog.Logger
}

// SendEmailHandler handles HTTP requests for sending an email.
//
// It is responsible for:
//   - Reading and validating the JSON payload
//   - Converting the DTO into an email message
//   - Delegating the request to the application use case
//   - Mapping infrastructure errors to HTTP responses
//
// Possible responses:
//   - 202 Accepted: request accepted for asynchronous processing
//   - 400 Bad Request: malformed or invalid JSON payload
//   - 400 Bad Request: type does not exists
//   - 422 Unprocessable Entity: field validation error
//   - 500 Internal Server Error: unexpected internal failure
func (s *SendEmailHandler) SendEmailHandler(w http.ResponseWriter, r *http.Request) {
	var dto MessageDTO

	err := s.readJSON(w, r, &dto)
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

	message, err := message.NewMessage(dto.Id, dto.To, dto.Type, dto.Variables)
	if err != nil {
		s.Logger.ErrorContext(
			r.Context(),
			"email validation error",
			"error", err,
			"method", r.Method,
			"path", r.URL.Path,
		)

		var fieldErr *apperrors.InvalidFieldError
		if errors.As(err, &fieldErr) {
			s.respond(
				w, r, http.StatusUnprocessableEntity,
				envelope{"error": fieldErr.Error(), "field": fieldErr.FieldName},
			)
			return
		}

		s.respond(
			w, r, http.StatusInternalServerError,
			envelope{"error": "internal server error"},
		)
		return
	}

	err = s.RequestSendEmailUsecase.Execute(*message)
	if err != nil {
		s.Logger.ErrorContext(
			r.Context(),
			"internal error while requesting email",
			"error", err,
			"method", r.Method,
			"path", r.URL.Path,
		)

		s.respond(
			w, r, http.StatusInternalServerError,
			envelope{"error": "internal server error"},
		)
		return
	}

	s.respond(w, r, http.StatusAccepted, envelope{"status": "accepted"})
}
