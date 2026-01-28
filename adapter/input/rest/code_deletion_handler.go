package rest

import (
	"net/http"
)

func (h *HandlerEmail) SendDeletionCodeHandler(w http.ResponseWriter, r *http.Request) {
	var dto DeletionCodeDTO
	h.handleEmailRequest(w, r, &dto)
}
