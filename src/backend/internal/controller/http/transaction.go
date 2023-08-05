package http

import (
	"APG6/internal/entity"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"path"
)

func (h *handler) AddTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	role, ok := r.Context().Value("role_id").(uint8)
	if !ok {
		WriteError(w, errors.New("invalid type of role_id"), http.StatusUnauthorized)
		return
	}

	store := entity.Transaction{}
	if err := json.NewDecoder(r.Body).Decode(&store); err != nil {
		WriteError(w, err, http.StatusBadRequest)
		return
	}

	if err := h.tableManager.AddTransaction(context.Background(), role, store); err != nil {
		WriteError(w, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *handler) GetTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	role, ok := r.Context().Value("role_id").(uint8)
	if !ok {
		WriteError(w, errors.New("invalid type of role_id"), http.StatusUnauthorized)
		return
	}

	transactions, err := h.tableManager.GetTransaction(context.Background(), role)
	if err != nil {
		WriteError(w, err, http.StatusInternalServerError)
		return
	}

	WriteResponseJson(w, transactions)
}

func (h *handler) UpdateTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	role, ok := r.Context().Value("role_id").(uint8)
	if !ok {
		WriteError(w, errors.New("invalid type of role_id"), http.StatusUnauthorized)
		return
	}

	transaction := entity.Transaction{}
	if err := json.NewDecoder(r.Body).Decode(&transaction); err != nil {
		WriteError(w, err, http.StatusBadRequest)
		return
	}

	if err := h.tableManager.UpdateTransaction(context.Background(), role, transaction); err != nil {
		WriteError(w, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *handler) DeleteTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	role, ok := r.Context().Value("role_id").(uint8)
	if !ok {
		WriteError(w, errors.New("invalid type of role_id"), http.StatusUnauthorized)
		return
	}

	if err := h.tableManager.DeleteTransaction(context.Background(), role, path.Base(r.URL.Path)); err != nil {
		WriteError(w, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
