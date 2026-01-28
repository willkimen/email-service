package rest

import (
	"net/http"
)

func (h *HandlerEmail) NotifyResetPasswordHandler(w http.ResponseWriter, r *http.Request) {
	var dto NotifyResetPasswordDTO
	h.handleEmailRequest(w, r, &dto)
}
