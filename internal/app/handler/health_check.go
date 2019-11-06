package handler

import (
	"net/http"

	"github.com/kitabisa/go-bootstrap/internal/pkg/commons"
)

// HealthCheck checking if all work well
func (h *Handler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	err, rc := h.services.HealthCheck.HealthCheck()

	if err != nil {
		h.WriteResponse(w, http.StatusInternalServerError, rc, nil, nil)
		return
	}

	h.WriteResponse(w, http.StatusOK, commons.RCSuccess, nil, nil)
	return
}
