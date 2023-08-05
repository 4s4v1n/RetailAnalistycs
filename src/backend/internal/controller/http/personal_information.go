package http

import (
	"APG6/internal/entity"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"path"
)

func (h *handler) AddPersonalInformation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	role, ok := r.Context().Value("role_id").(uint8)
	if !ok {
		WriteError(w, errors.New("invalid type of role_id"), http.StatusUnauthorized)
		return
	}

	information := entity.PersonalInformation{}
	if err := json.NewDecoder(r.Body).Decode(&information); err != nil {
		WriteError(w, err, http.StatusBadRequest)
		return
	}

	if err := h.tableManager.AddPersonalInformation(context.Background(), role, information); err != nil {
		WriteError(w, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *handler) GetPersonalInformation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	role, ok := r.Context().Value("role_id").(uint8)
	if !ok {
		WriteError(w, errors.New("invalid type of role_id"), http.StatusUnauthorized)
		return
	}

	information, err := h.tableManager.GetPersonalInformation(context.Background(), role)
	if err != nil {
		WriteError(w, err, http.StatusInternalServerError)
		return
	}

	WriteResponseJson(w, information)
}

func (h *handler) UpdatePersonalInformation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	role, ok := r.Context().Value("role_id").(uint8)
	if !ok {
		WriteError(w, errors.New("invalid type of role_id"), http.StatusUnauthorized)
		return
	}

	information := entity.PersonalInformation{}
	if err := json.NewDecoder(r.Body).Decode(&information); err != nil {
		WriteError(w, err, http.StatusBadRequest)
		return
	}

	if err := h.tableManager.UpdatePersonalInformation(context.Background(), role, information); err != nil {
		WriteError(w, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *handler) DeletePersonalInformation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	role, ok := r.Context().Value("role_id").(uint8)
	if !ok {
		WriteError(w, errors.New("invalid type of role_id"), http.StatusUnauthorized)
		return
	}

	if err := h.tableManager.DeletePersonalInformation(context.Background(), role, path.Base(r.URL.Path)); err != nil {
		WriteError(w, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
