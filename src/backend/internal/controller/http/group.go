package http

import (
	"context"
	"errors"
	"net/http"
)

func (h *handler) GetGroups(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	role, ok := r.Context().Value("role_id").(uint8)
	if !ok {
		WriteError(w, errors.New("invalid type of role_id"), http.StatusUnauthorized)
		return
	}

	groups, err := h.viewManager.GetGroups(context.Background(), role)
	if err != nil {
		WriteError(w, err, http.StatusInternalServerError)
		return
	}

	WriteResponseJson(w, groups)
}
