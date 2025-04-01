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

func NewAuthRouter(service *services.AuthService) *AuthHandler {
	return &AuthHandler{auth: service}
}

func (h *AuthHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/auth/login", h.login)
	mux.HandleFunc("/auth/signup", h.signup)
}

func (h *AuthHandler) login(w http.ResponseWriter, r *http.Request) {
	cmd := &commands.LoginCommand{}
	ctx, err := request.ParseRequestBody(r, cmd)
	if err != nil {
		response.WriteError(w, err)
		return
	}

	if token, err := h.auth.Login(ctx, cmd); err != nil {
		response.WriteError(w, err)
		return
	} else {
		response.WriteJSON(w, http.StatusOK, token)
	}

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
