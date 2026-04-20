package rest

import (
	"net/http"
)

// NotifyEmailVerificationHandler handles requests to send an email verification notification email.
//
// It decodes the incoming JSON payload into a NotifyEmailVerificationDTO and delegates
// the request handling to the shared email request helper.
func (s *SendEmailHandler) NotifyEmailVerificationHandler(w http.ResponseWriter, r *http.Request) {
	var dto NotifyEmailVerificationDTO
	s.handleEmailRequest(w, r, &dto)
}

// NotifyChangeEmailHandler handles requests to send a change email notification.
//
// It parses the request body into a NotifyChangeEmailDTO and forwards it to the
// shared email request handler for validation and processing.
func (s *SendEmailHandler) NotifyChangeEmailHandler(w http.ResponseWriter, r *http.Request) {
	var dto NotifyChangeEmailDTO
	s.handleEmailRequest(w, r, &dto)
}

// NotifyChangePasswordHandler handles requests to send a change password notification.
//
// It decodes the request body into a NotifyChangePasswordDTO and delegates
// validation and processing to the shared email request handler.
func (s *SendEmailHandler) NotifyChangePasswordHandler(w http.ResponseWriter, r *http.Request) {
	var dto NotifyChangePasswordDTO
	s.handleEmailRequest(w, r, &dto)
}

// NotifyDeletionHandler handles requests to send an account deletion notification.
//
// It decodes the incoming request into a NotifyDeletionDTO and forwards it
// to the common email request handling flow, which performs validation and
// dispatches the email request asynchronously.
func (s *SendEmailHandler) NotifyDeletionHandler(w http.ResponseWriter, r *http.Request) {
	var dto NotifyDeletionDTO
	s.handleEmailRequest(w, r, &dto)
}

// NotifyResetPasswordHandler handles requests to notify the user that a password
// reset has occurred.
//
// It parses the request payload into a NotifyResetPasswordDTO and delegates
// processing to the shared email request handler, which validates the data
// and triggers the email request asynchronously.
func (s *SendEmailHandler) NotifyResetPasswordHandler(w http.ResponseWriter, r *http.Request) {
	var dto NotifyResetPasswordDTO
	s.handleEmailRequest(w, r, &dto)
}
