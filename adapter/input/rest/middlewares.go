package rest

import "net/http"

func (s *SendEmailHandler) recoverPanicMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {

				s.Logger.ErrorContext(
					r.Context(),
					"panic recovered in http handler",
					"panic", err,
					"method", r.Method,
					"path", r.URL.Path,
				)

				w.Header().Set("Connection", "close")
				s.serverErrorResponse(w, r)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
