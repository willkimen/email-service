package rest

import (
	"net/http"
)

func (h *HandlerEmail) SendResetPasswordCodeHandler(w http.ResponseWriter, r *http.Request) {
	var dto ResetPasswordCodeDTO
	h.handleEmailRequest(w, r, &dto)
}
