package rest

import (
	"net/http"
)

func (h *HandlerEmail) NotifyChangeEmailHandler(w http.ResponseWriter, r *http.Request) {
	var dto NotifyChangeEmailDTO
	h.handleEmailRequest(w, r, &dto)
}
