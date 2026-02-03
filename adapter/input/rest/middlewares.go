package rest

import "net/http"

func (s *SendEmailHandler) recoverPanicMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				w.Header().Set("Connection", "close")
				s.serverErrorResponse(w, r)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
