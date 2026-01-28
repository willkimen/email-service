package rest

import (
	"net/http"
)

func (h *HandlerEmail) SendChangePasswordCodeHandler(w http.ResponseWriter, r *http.Request) {
	var dto ChangePasswordCodeDTO
	h.handleEmailRequest(w, r, &dto)
}
