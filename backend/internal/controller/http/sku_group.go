package http

import (
	"APG6/internal/entity"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"path"
)

func (h *handler) AddSkuGroup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	role, ok := r.Context().Value("role_id").(uint8)
	if !ok {
		WriteError(w, errors.New("invalid type of role_id"), http.StatusUnauthorized)
		return
	}

	group := entity.SkuGroup{}
	if err := json.NewDecoder(r.Body).Decode(&group); err != nil {
		WriteError(w, err, http.StatusBadRequest)
		return
	}

	if err := h.tableManager.AddSkuGroup(context.Background(), role, group); err != nil {
		WriteError(w, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *handler) GetSkuGroup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	role, ok := r.Context().Value("role_id").(uint8)
	if !ok {
		WriteError(w, errors.New("invalid type of role_id"), http.StatusUnauthorized)
		return
	}

	group, err := h.tableManager.GetSkuGroup(context.Background(), role)
	if err != nil {
		WriteError(w, err, http.StatusInternalServerError)
		return
	}

	WriteResponseJson(w, group)
}

func (h *handler) UpdateSkuGroup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	role, ok := r.Context().Value("role_id").(uint8)
	if !ok {
		WriteError(w, errors.New("invalid type of role_id"), http.StatusUnauthorized)
		return
	}

	group := entity.SkuGroup{}
	if err := json.NewDecoder(r.Body).Decode(&group); err != nil {
		WriteError(w, err, http.StatusBadRequest)
		return
	}

	if err := h.tableManager.UpdateSkuGroup(context.Background(), role, group); err != nil {
		WriteError(w, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *handler) DeleteSkuGroup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	role, ok := r.Context().Value("role_id").(uint8)
	if !ok {
		WriteError(w, errors.New("invalid type of role_id"), http.StatusUnauthorized)
		return
	}

	if err := h.tableManager.DeleteSkuGroup(context.Background(), role, path.Base(r.URL.Path)); err != nil {
		WriteError(w, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
