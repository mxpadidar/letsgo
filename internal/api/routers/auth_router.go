package routers

import (
	"encoding/json"
	"net/http"

	"github.com/mxpadidar/letsgo/internal/api/response"
	"github.com/mxpadidar/letsgo/internal/domain/commands"
	"github.com/mxpadidar/letsgo/internal/domain/errors"
	"github.com/mxpadidar/letsgo/internal/domain/stores"
)

type AuthRouter struct {
	mux       *http.ServeMux
	userStore stores.UserStore
}

func NewAuthRouter(mux *http.ServeMux, userStore stores.UserStore) *AuthRouter {
	return &AuthRouter{mux: mux, userStore: userStore}
}

func (router *AuthRouter) Load() {
	router.mux.HandleFunc("POST /auth/signup", router.signup)
}

func (router *AuthRouter) signup(w http.ResponseWriter, r *http.Request) {
	cmd := &commands.SignupCmd{}
	if err := json.NewDecoder(r.Body).Decode(cmd); err != nil {
		response.WriteError(w, errors.NewErr(errors.ErrValidation, "invalid request body", err))
		return
	}

	user, err := cmd.Execute(r.Context(), router.userStore)
	if err != nil {
		response.WriteError(w, err)
		return
	}

	response.WriteOk(w, "signup successful", user, http.StatusCreated)
}
