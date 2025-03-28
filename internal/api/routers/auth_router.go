package routers

import (
	"encoding/json"
	"net/http"

	"github.com/mxpadidar/letsgo/internal/api/helpers"
	"github.com/mxpadidar/letsgo/internal/domain/commands"
	"github.com/mxpadidar/letsgo/internal/domain/errors"
	"github.com/mxpadidar/letsgo/internal/domain/services"
	"github.com/mxpadidar/letsgo/internal/domain/stores"
)

type AuthRouter struct {
	mux          *http.ServeMux
	userStore    stores.UserStore
	tokenServ    services.TokenService
	passwordServ services.PasswordService
}

func NewAuthRouter(
	mux *http.ServeMux,
	userStore stores.UserStore,
	tokenServ services.TokenService,
	passwordServ services.PasswordService,
) *AuthRouter {
	return &AuthRouter{
		mux:          mux,
		userStore:    userStore,
		tokenServ:    tokenServ,
		passwordServ: passwordServ,
	}
}

func (router *AuthRouter) Load() {
	router.mux.HandleFunc("POST /auth/signup", router.signup)
	router.mux.HandleFunc("POST /auth/login", router.login)
}

func (router *AuthRouter) signup(w http.ResponseWriter, r *http.Request) {
	cmd := &commands.SignupCmd{}
	if err := json.NewDecoder(r.Body).Decode(cmd); err != nil {
		helpers.WriteError(w, errors.NewErr(errors.ErrValidation, "invalid request body", err))
		return
	}

	user, err := cmd.Execute(r.Context(), router.userStore, router.passwordServ)
	if err != nil {
		helpers.WriteError(w, err)
		return
	}

	helpers.WriteJSON(w, user, http.StatusCreated)
}

func (router *AuthRouter) login(w http.ResponseWriter, r *http.Request) {
	cmd := &commands.AuthCmd{}
	if err := json.NewDecoder(r.Body).Decode(cmd); err != nil {
		helpers.WriteError(w, errors.NewErr(errors.ErrValidation, "invalid request body", err))
		return
	}

	// get authenticated user
	user, err := cmd.Execute(r.Context(), router.userStore, router.passwordServ)
	if err != nil {
		helpers.WriteError(w, err)
		return
	}

	// generate token for authenticated user
	token, err := router.tokenServ.Encode(r.Context(), user)
	if err != nil {
		helpers.WriteError(w, err)
		return
	}

	helpers.WriteJSON(w, token, http.StatusOK)
}
