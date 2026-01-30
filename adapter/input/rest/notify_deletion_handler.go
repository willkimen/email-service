package rest

import (
	"net/http"
)

// NotifyDeletionHandler handles requests to send an account deletion notification.
//
// It decodes the incoming request into a NotifyDeletionDTO and forwards it
// to the common email request handling flow, which performs validation and
// dispatches the email request asynchronously.
func (h *HandlerEmail) NotifyDeletionHandler(w http.ResponseWriter, r *http.Request) {
	var dto NotifyDeletionDTO
	h.handleEmailRequest(w, r, &dto)
}
