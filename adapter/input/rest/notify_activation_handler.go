package rest

import (
	"net/http"
)

func (h *HandlerEmail) NotifyActivationHandler(w http.ResponseWriter, r *http.Request) {
	var dto NotifyActivationDTO
	h.handleEmailRequest(w, r, &dto)
}
