package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/mxpadidar/letsgo/internal/core/commands"
	"github.com/mxpadidar/letsgo/internal/core/services"
	"github.com/mxpadidar/letsgo/internal/core/stores"
)

type AuthHandler struct {
	mux         *http.ServeMux
	userStore   stores.UserStore
	hashService services.HashService
}

func NewAuthHandler(mux *http.ServeMux, userStore stores.UserStore, hashService services.HashService) *AuthHandler {
	return &AuthHandler{mux: mux, userStore: userStore, hashService: hashService}
}

func (h *AuthHandler) SetupRoutes() {
	h.mux.HandleFunc("POST /auth/login", h.loginHandler)
	h.mux.HandleFunc("POST /auth/register", h.registerHandler)
}

func (h *AuthHandler) loginHandler(w http.ResponseWriter, r *http.Request) {
	authCmd := commands.NewAuthCmd(h.userStore, h.hashService)

	if err := json.NewDecoder(r.Body).Decode(&authCmd); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := authCmd.Execute(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func (h *AuthHandler) registerHandler(w http.ResponseWriter, r *http.Request) {
	cmd := commands.NewRegisterCmd(h.userStore, h.hashService)
	err := json.NewDecoder(r.Body).Decode(&cmd)
	println(cmd.Username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := cmd.Execute(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}
