package rest

import (
	"net/http"
)

// SendResetPasswordCodeHandler handles HTTP requests for sending
// a reset password verification code.
//
// This handler is responsible only for wiring the HTTP layer.
// The full request lifecycle (JSON parsing, validation, use case
// execution and HTTP response mapping) is delegated to the shared helper.
func (s *SendEmailHandler) SendResetPasswordCodeHandler(w http.ResponseWriter, r *http.Request) {
	// ResetPasswordCodeDTO represents the expected payload
	// for requesting a password reset verification code.
	var dto ResetPasswordCodeDTO

	// Delegate the request handling flow to the shared helper.
	s.handleEmailRequest(w, r, &dto)
}
