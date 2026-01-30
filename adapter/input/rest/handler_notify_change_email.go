package rest

import (
	"net/http"
)

// NotifyChangeEmailHandler handles requests to send a change email notification.
//
// It parses the request body into a NotifyChangeEmailDTO and forwards it to the
// shared email request handler for validation and processing.
func (h *SendEmailHandler) NotifyChangeEmailHandler(w http.ResponseWriter, r *http.Request) {
	var dto NotifyChangeEmailDTO
	h.handleEmailRequest(w, r, &dto)
}
