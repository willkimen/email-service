package rest

import (
	"net/http"
)

func (h *HandlerEmail) NotifyDeletionHandler(w http.ResponseWriter, r *http.Request) {
	var dto NotifyDeletionDTO
	h.handleEmailRequest(w, r, &dto)
}
