package rest

import (
	"net/http"
)

// SendChangeEmailCodeHandler handles HTTP requests for sending
// a change email verification code.
//
// This handler is limited to HTTP concerns and delegates
// request parsing, validation, use case execution and
// error-to-response mapping to the shared helper.
func (s *SendEmailHandler) SendChangeEmailCodeHandler(w http.ResponseWriter, r *http.Request) {
	// ChangeEmailCodeDTO represents the expected payload
	// for requesting a change email verification code.
	var dto ChangeEmailCodeDTO

	// Delegate the request handling flow to the shared helper.
	s.handleEmailRequest(w, r, &dto)
}
