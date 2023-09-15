package http

import (
	"context"
	"errors"
	"net/http"
	"path"
)

func (h *handler) Import(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	role, ok := r.Context().Value("role_id").(uint8)
	if !ok {
		WriteError(w, errors.New("invalid type of role_id"), http.StatusUnauthorized)
		return
	}

	if err := h.dataManager.Import(context.Background(), role, path.Base(r.URL.Path), r.Body); err != nil {
		WriteError(w, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *handler) Export(w http.ResponseWriter, r *http.Request) {
	role, ok := r.Context().Value("role_id").(uint8)
	if !ok {
		WriteError(w, errors.New("invalid type of role_id"), http.StatusUnauthorized)
		return
	}

	data, err := h.dataManager.Export(context.Background(), role, path.Base(r.URL.Path))
	if err != nil {
		WriteError(w, err, http.StatusInternalServerError)
		return
	}

	WriteResponseCsv(w, data)
}
