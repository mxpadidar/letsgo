package handlers

import (
	"net/http"

	"github.com/mxpadidar/letsgo/internal/api/request"
	"github.com/mxpadidar/letsgo/internal/api/response"
	"github.com/mxpadidar/letsgo/internal/domain/commands"
	"github.com/mxpadidar/letsgo/internal/domain/services"
)

type AuthHandler struct {
	auth *services.AuthService
}

func NewAuthHandler(service *services.AuthService) *AuthHandler {
	return &AuthHandler{auth: service}
}

func (h *AuthHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /auth/signup", h.signup)
	mux.HandleFunc("POST /auth/tokens/issue", h.issueTokens)
	mux.HandleFunc("POST /auth/tokens/refresh", h.refreshTokens)
	mux.HandleFunc("DELETE /auth/tokens/revoke", h.revokeTokens)
}

func (h *AuthHandler) signup(w http.ResponseWriter, r *http.Request) {
	cmd := &commands.SignupCommand{}
	ctx, err := request.ParseRequestBody(r, cmd)
	if err != nil {
		response.WriteError(w, err)
		return
	}

	if user, err := h.auth.Signup(ctx, cmd); err != nil {
		response.WriteError(w, err)
		return
	} else {
		response.WriteJSON(w, http.StatusOK, user)
	}
}

func (h *AuthHandler) issueTokens(w http.ResponseWriter, r *http.Request) {
	cmd := &commands.CreatePermitCmd{}
	ctx, err := request.ParseRequestBody(r, cmd)
	if err != nil {
		response.WriteError(w, err)
		return
	}

	if token, err := h.auth.CreatePermit(ctx, cmd); err != nil {
		response.WriteError(w, err)
		return
	} else {
		response.WriteJSON(w, http.StatusOK, token)
	}
}

func (h *AuthHandler) refreshTokens(w http.ResponseWriter, r *http.Request) {
	cmd := &commands.RotatePermitCmd{}
	ctx, err := request.ParseRequestBody(r, cmd)
	if err != nil {
		response.WriteError(w, err)
		return
	}

	if token, err := h.auth.RotatePermit(ctx, cmd); err != nil {
		response.WriteError(w, err)
		return
	} else {
		response.WriteJSON(w, http.StatusOK, token)
	}
}

func (h *AuthHandler) revokeTokens(w http.ResponseWriter, r *http.Request) {
	if err := h.auth.RevokePermit(r.Context()); err != nil {
		response.WriteError(w, err)
		return
	}

	response.WriteJSON(w, http.StatusNoContent, nil)
}
