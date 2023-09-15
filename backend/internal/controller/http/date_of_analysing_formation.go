package http

import (
	"APG6/internal/entity"
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

func (h *handler) GetDateOfAnalysingFormation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	role, ok := r.Context().Value("role_id").(uint8)
	if !ok {
		WriteError(w, errors.New("invalid type of role_id"), http.StatusUnauthorized)
		return
	}

	formation, err := h.tableManager.GetDateOfAnalysingFormation(context.Background(), role)
	if err != nil {
		WriteError(w, err, http.StatusInternalServerError)
		return
	}

	WriteResponseJson(w, formation)
}

func (h *handler) UpdateDateOfAnalysingFormation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	role, ok := r.Context().Value("role_id").(uint8)
	if !ok {
		WriteError(w, errors.New("invalid type of role_id"), http.StatusUnauthorized)
		return
	}

	formation := entity.UpdateFormation{}
	if err := json.NewDecoder(r.Body).Decode(&formation); err != nil {
		WriteError(w, err, http.StatusBadRequest)
		return
	}

	if err := h.tableManager.UpdateDateOfAnalysingFormation(context.Background(), role, formation); err != nil {
		WriteError(w, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
