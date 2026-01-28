package rest

import (
	"net/http"
)

func (h *HandlerEmail) NotifyChangePasswordHandler(w http.ResponseWriter, r *http.Request) {
	var dto NotifyChangePasswordDTO
	h.handleEmailRequest(w, r, &dto)
}
