package rest

import (
	"net/http"
)

func (h *HandlerEmail) SendActivationCodeHandler(w http.ResponseWriter, r *http.Request) {
	var dto ActivationCodeDTO
	h.handleEmailRequest(w, r, &dto)
}
