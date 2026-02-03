package rest

import (
	"net/http"
)

// NotifyChangePasswordHandler handles requests to send a change password notification.
//
// It decodes the request body into a NotifyChangePasswordDTO and delegates
// validation and processing to the shared email request handler.
func (s *SendEmailHandler) NotifyChangePasswordHandler(w http.ResponseWriter, r *http.Request) {
	var dto NotifyChangePasswordDTO
	s.handleEmailRequest(w, r, &dto)
}
