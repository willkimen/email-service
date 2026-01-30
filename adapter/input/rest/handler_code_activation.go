package rest

import (
	"net/http"
)

// SendActivationCodeHandler handles HTTP requests for sending
// an activation code email.
//
// This handler is responsible only for HTTP-level concerns.
// It delegates JSON parsing, validation, domain interaction
// and response mapping to the shared helper.
func (h *SendEmailHandler) SendActivationCodeHandler(w http.ResponseWriter, r *http.Request) {
	// ActivationCodeDTO represents the expected payload
	// for requesting an activation code email.
	var dto ActivationCodeDTO

	// Delegate the full request lifecycle to the shared helper:
	// - read and validate the request body
	// - convert DTO to a domain message
	// - invoke the use case
	// - map domain and infrastructure errors to HTTP responses
	h.handleEmailRequest(w, r, &dto)
}
