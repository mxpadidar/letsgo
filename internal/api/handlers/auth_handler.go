package handlers

import (
	"net/http"

	"github.com/mxpadidar/letsgo/internal/api/request"
	"github.com/mxpadidar/letsgo/internal/api/response"
	"github.com/mxpadidar/letsgo/internal/core/commands"
	"github.com/mxpadidar/letsgo/internal/core/services"
)

type AuthHandler struct {
	service *services.AuthService
}

func NewAuthHandler(service *services.AuthService) *AuthHandler {
	return &AuthHandler{service: service}
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
		response.WriteError(w, h.service.Logger, err)
		return
	}

	if user, err := h.service.Signup(ctx, cmd); err != nil {
		response.WriteError(w, h.service.Logger, err)
		return
	} else {
		response.WriteJSON(w, h.service.Logger, http.StatusOK, user)
	}
}

func (h *AuthHandler) issueTokens(w http.ResponseWriter, r *http.Request) {
	cmd := &commands.CreatePermitCmd{}
	ctx, err := request.ParseRequestBody(r, cmd)
	if err != nil {
		response.WriteError(w, h.service.Logger, err)
		return
	}

	if token, err := h.service.CreatePermit(ctx, cmd); err != nil {
		response.WriteError(w, h.service.Logger, err)
		return
	} else {
		response.WriteJSON(w, h.service.Logger, http.StatusOK, token)
	}
}

func (h *AuthHandler) refreshTokens(w http.ResponseWriter, r *http.Request) {
	cmd := &commands.RotatePermitCmd{}
	ctx, err := request.ParseRequestBody(r, cmd)
	if err != nil {
		response.WriteError(w, h.service.Logger, err)
		return
	}

	if token, err := h.service.RotatePermit(ctx, cmd); err != nil {
		response.WriteError(w, h.service.Logger, err)
		return
	} else {
		response.WriteJSON(w, h.service.Logger, http.StatusOK, token)
	}
}

func (h *AuthHandler) revokeTokens(w http.ResponseWriter, r *http.Request) {
	if err := h.service.RevokePermit(r.Context()); err != nil {
		response.WriteError(w, h.service.Logger, err)
		return
	}

	response.WriteJSON(w, h.service.Logger, http.StatusNoContent, nil)
}
