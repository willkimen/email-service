package rest

import (
	"net/http"
)

// NotifyResetPasswordHandler handles requests to notify the user that a password
// reset has occurred.
//
// It parses the request payload into a NotifyResetPasswordDTO and delegates
// processing to the shared email request handler, which validates the data
// and triggers the email request asynchronously.
func (h *HandlerEmail) NotifyResetPasswordHandler(w http.ResponseWriter, r *http.Request) {
	var dto NotifyResetPasswordDTO
	h.handleEmailRequest(w, r, &dto)
}
