package http

import (
	"APG6/internal/controller/auth"
	"context"
	"encoding/json"
	"github.com/go-chi/jwtauth"
	"net/http"
)

const (
	_roleId = "role_id"
)

type authorizationRequest struct {
	Role     string `json:"role"`
	Password string `json:"password"`
}

type refreshRequest struct {
	Token string `json:"token"`
}

type authorizationResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	Role         string `json:"role"`
}

func Authorizer() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			id, err := auth.DecodeToken(jwtauth.TokenFromHeader(r))
			if err != nil {
				WriteError(w, err, http.StatusUnauthorized)
				return
			}
			next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), _roleId, id)))
		})
	}
}

func (h *handler) Authorization(w http.ResponseWriter, r *http.Request) {
	defer func() {
		_ = r.Body.Close()
	}()

	var req authorizationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		WriteError(w, err, http.StatusBadRequest)
		return
	}

	token, err := auth.EncodeToken(req.Role, req.Password)
	if err != nil {
		WriteError(w, err, http.StatusUnauthorized)
		return
	}
	refresh, err := auth.RefreshToken()
	if err != nil {
		WriteError(w, err, http.StatusUnauthorized)
		return
	}
	if err = auth.SaveSession(refresh, req.Role, req.Password); err != nil {
		WriteError(w, err, http.StatusUnauthorized)
		return
	}
	WriteResponseJson(w, authorizationResponse{
		AccessToken:  token,
		RefreshToken: refresh,
		Role:         req.Role,
	})
}

func (h *handler) Refresh(w http.ResponseWriter, r *http.Request) {
	defer func() {
		_ = r.Body.Close()
	}()

	var req refreshRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		WriteError(w, err, http.StatusBadRequest)
		return
	}
	login, pass, err := auth.GetSession(req.Token)
	if err != nil {
		WriteError(w, err, http.StatusUnauthorized)
		return
	}
	token, err := auth.EncodeToken(login, pass)
	if err != nil {
		WriteError(w, err, http.StatusUnauthorized)
		return
	}
	refresh, err := auth.RefreshToken()
	if err != nil {
		WriteError(w, err, http.StatusUnauthorized)
		return
	}
	if err = auth.SaveSession(refresh, login, pass); err != nil {
		WriteError(w, err, http.StatusUnauthorized)
		return
	}
	WriteResponseJson(w, authorizationResponse{
		AccessToken:  token,
		RefreshToken: refresh,
		Role:         login,
	})
}
