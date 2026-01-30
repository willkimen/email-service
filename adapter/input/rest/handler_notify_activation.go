package rest

import (
	"net/http"
)

// NotifyActivationHandler handles requests to send an activation notification email.
//
// It decodes the incoming JSON payload into a NotifyActivationDTO and delegates
// the request handling to the shared email request helper.
func (h *SendEmailHandler) NotifyActivationHandler(w http.ResponseWriter, r *http.Request) {
	var dto NotifyActivationDTO
	h.handleEmailRequest(w, r, &dto)
}
