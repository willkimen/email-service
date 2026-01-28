package rest

import (
	"net/http"
)

func (h *HandlerEmail) SendChangeEmailCodeHandler(w http.ResponseWriter, r *http.Request) {
	var dto ChangeEmailCodeDTO
	h.handleEmailRequest(w, r, &dto)
}
