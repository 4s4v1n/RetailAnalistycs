package http

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"path"

	"APG6/internal/entity"
)

func (h *handler) AddCard(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	role, ok := r.Context().Value("role_id").(uint8)
	if !ok {
		WriteError(w, errors.New("invalid type of role_id"), http.StatusUnauthorized)
		return
	}

	card := entity.Card{}
	if err := json.NewDecoder(r.Body).Decode(&card); err != nil {
		WriteError(w, err, http.StatusBadRequest)
		return
	}

	if err := h.tableManager.AddCard(context.Background(), role, card); err != nil {
		WriteError(w, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *handler) GetCard(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	role, ok := r.Context().Value("role_id").(uint8)
	if !ok {
		WriteError(w, errors.New("invalid type of role_id"), http.StatusUnauthorized)
		return
	}

	cards, err := h.tableManager.GetCard(context.Background(), role)
	if err != nil {
		WriteError(w, err, http.StatusInternalServerError)
		return
	}

	WriteResponseJson(w, cards)
}

func (h *handler) UpdateCard(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	role, ok := r.Context().Value("role_id").(uint8)
	if !ok {
		WriteError(w, errors.New("invalid type of role_id"), http.StatusUnauthorized)
		return
	}

	card := entity.Card{}
	if err := json.NewDecoder(r.Body).Decode(&card); err != nil {
		WriteError(w, err, http.StatusBadRequest)
		return
	}

	if err := h.tableManager.UpdateCard(context.Background(), role, card); err != nil {
		WriteError(w, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *handler) DeleteCard(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	role, ok := r.Context().Value("role_id").(uint8)
	if !ok {
		WriteError(w, errors.New("invalid type of role_id"), http.StatusUnauthorized)
		return
	}

	if err := h.tableManager.DeleteCard(context.Background(), role, path.Base(r.URL.Path)); err != nil {
		WriteError(w, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
