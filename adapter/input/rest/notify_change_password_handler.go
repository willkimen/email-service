package rest

import (
	"net/http"
)

// NotifyChangePasswordHandler handles requests to send a change password notification.
//
// It decodes the request body into a NotifyChangePasswordDTO and delegates
// validation and processing to the shared email request handler.
func (h *HandlerEmail) NotifyChangePasswordHandler(w http.ResponseWriter, r *http.Request) {
	var dto NotifyChangePasswordDTO
	h.handleEmailRequest(w, r, &dto)
}
