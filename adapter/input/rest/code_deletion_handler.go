package rest

import (
	"net/http"
)

// SendDeletionCodeHandler handles HTTP requests for sending
// an account deletion verification code.
//
// This handler only coordinates the HTTP layer and delegates
// request parsing, validation, use case execution and response
// handling to the shared helper.
func (h *HandlerEmail) SendDeletionCodeHandler(w http.ResponseWriter, r *http.Request) {
	// DeletionCodeDTO represents the expected payload
	// for requesting an account deletion verification code.
	var dto DeletionCodeDTO

	// Delegate the request handling flow to the shared helper.
	h.handleEmailRequest(w, r, &dto)
}
