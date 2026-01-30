package rest

import (
	"net/http"
)

// SendChangePasswordCodeHandler handles HTTP requests for sending
// a change password verification code.
//
// This handler is responsible only for HTTP-level concerns and
// delegates request decoding, validation, use case execution
// and response handling to the shared helper.
func (h *HandlerEmail) SendChangePasswordCodeHandler(w http.ResponseWriter, r *http.Request) {
	// ChangePasswordCodeDTO represents the expected payload
	// for requesting a change password verification code.
	var dto ChangePasswordCodeDTO

	// Delegate the full request handling flow to the shared helper.
	h.handleEmailRequest(w, r, &dto)
}
