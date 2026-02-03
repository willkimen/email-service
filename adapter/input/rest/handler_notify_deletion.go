package rest

import (
	"net/http"
)

// NotifyDeletionHandler handles requests to send an account deletion notification.
//
// It decodes the incoming request into a NotifyDeletionDTO and forwards it
// to the common email request handling flow, which performs validation and
// dispatches the email request asynchronously.
func (s *SendEmailHandler) NotifyDeletionHandler(w http.ResponseWriter, r *http.Request) {
	var dto NotifyDeletionDTO
	s.handleEmailRequest(w, r, &dto)
}
