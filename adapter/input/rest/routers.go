package rest

import "net/http"

func (s *SendEmailHandler) Routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /api/v1/email/activation/code", s.SendActivationCodeHandler)
	mux.HandleFunc("POST /api/v1/email/activation/notify", s.NotifyActivationHandler)

	mux.HandleFunc("POST /api/v1/email/change-email/code", s.SendChangeEmailCodeHandler)
	mux.HandleFunc("POST /api/v1/email/change-email/notify", s.NotifyChangeEmailHandler)

	mux.HandleFunc("POST /api/v1/email/change-password/code", s.SendChangePasswordCodeHandler)
	mux.HandleFunc("POST /api/v1/email/change-password/notify", s.NotifyChangePasswordHandler)

	mux.HandleFunc("POST /api/v1/email/reset-password/code", s.SendResetPasswordCodeHandler)
	mux.HandleFunc("POST /api/v1/email/reset-password/notify", s.NotifyResetPasswordHandler)

	mux.HandleFunc("POST /api/v1/email/deletion/code", s.SendDeletionCodeHandler)
	mux.HandleFunc("POST /api/v1/email/deletion/notify", s.NotifyDeletionHandler)

	return s.recoverPanicMiddleware(mux)
}
