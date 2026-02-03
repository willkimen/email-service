package rest

import (
	"net/http"
)

// NotifyActivationHandler handles requests to send an activation notification email.
//
// It decodes the incoming JSON payload into a NotifyActivationDTO and delegates
// the request handling to the shared email request helper.
func (s *SendEmailHandler) NotifyActivationHandler(w http.ResponseWriter, r *http.Request) {
	var dto NotifyActivationDTO
	s.handleEmailRequest(w, r, &dto)
}
