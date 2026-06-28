package rest

import "net/http"

func (s *SendEmailHandler) Routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /api/v1/email", s.SendEmailHandler)
	return s.recoverPanicMiddleware(mux)
}
