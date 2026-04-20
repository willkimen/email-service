package rest

import (
	"net/http"
)

// SendEmailVerificationCodeHandler handles HTTP requests for sending
// an email verification code email.
//
// This handler is responsible only for HTTP-level concerns.
// It delegates JSON parsing, validation, interaction
// and response mapping to the shared helper.
func (s *SendEmailHandler) SendEmailVerificationCodeHandler(w http.ResponseWriter, r *http.Request) {
	// EmailVerificationCodeDTO represents the expected payload
	// for requesting an email verification code email.
	var dto EmailVerificationCodeDTO

	// Delegate the full request lifecycle to the shared helper:
	// - read and validate the request body
	// - convert DTO to a email message
	// - invoke the use case
	// - map application and infrastructure errors to HTTP responses
	s.handleEmailRequest(w, r, &dto)
}

// SendChangeEmailCodeHandler handles HTTP requests for sending
// a change email verification code.
//
// This handler is limited to HTTP concerns and delegates
// request parsing, validation, use case execution and
// error-to-response mapping to the shared helper.
func (s *SendEmailHandler) SendChangeEmailCodeHandler(w http.ResponseWriter, r *http.Request) {
	// ChangeEmailCodeDTO represents the expected payload
	// for requesting a change email verification code.
	var dto ChangeEmailCodeDTO

	// Delegate the request handling flow to the shared helper.
	s.handleEmailRequest(w, r, &dto)
}

// SendChangePasswordCodeHandler handles HTTP requests for sending
// a change password verification code.
//
// This handler is responsible only for HTTP-level concerns and
// delegates request decoding, validation, use case execution
// and response handling to the shared helper.
func (s *SendEmailHandler) SendChangePasswordCodeHandler(w http.ResponseWriter, r *http.Request) {
	// ChangePasswordCodeDTO represents the expected payload
	// for requesting a change password verification code.
	var dto ChangePasswordCodeDTO

	// Delegate the full request handling flow to the shared helper.
	s.handleEmailRequest(w, r, &dto)
}

// SendDeletionCodeHandler handles HTTP requests for sending
// an account deletion verification code.
//
// This handler only coordinates the HTTP layer and delegates
// request parsing, validation, use case execution and response
// handling to the shared helper.
func (s *SendEmailHandler) SendDeletionCodeHandler(w http.ResponseWriter, r *http.Request) {
	// DeletionCodeDTO represents the expected payload
	// for requesting an account deletion verification code.
	var dto DeletionCodeDTO

	// Delegate the request handling flow to the shared helper.
	s.handleEmailRequest(w, r, &dto)
}

// SendResetPasswordCodeHandler handles HTTP requests for sending
// a reset password verification code.
//
// This handler is responsible only for wiring the HTTP layer.
// The full request lifecycle (JSON parsing, validation, use case
// execution and HTTP response mapping) is delegated to the shared helper.
func (s *SendEmailHandler) SendResetPasswordCodeHandler(w http.ResponseWriter, r *http.Request) {
	// ResetPasswordCodeDTO represents the expected payload
	// for requesting a password reset verification code.
	var dto ResetPasswordCodeDTO

	// Delegate the request handling flow to the shared helper.
	s.handleEmailRequest(w, r, &dto)
}
