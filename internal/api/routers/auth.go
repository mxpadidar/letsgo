package routers

import (
	"encoding/json"
	"net/http"

	"github.com/mxpadidar/letsgo/internal/api/response"
	"github.com/mxpadidar/letsgo/internal/core/commands"
	"github.com/mxpadidar/letsgo/internal/core/specs"
	"github.com/mxpadidar/letsgo/internal/core/stores"
	"github.com/mxpadidar/letsgo/internal/services/handlers"
)

type AuthRouter struct {
	mux          *http.ServeMux
	userStore    stores.UserStore
	tokenService specs.TokenService
	passService  specs.PasswordService
}

func NewAuthRouter(
	mux *http.ServeMux,
	userStore stores.UserStore,
	tokenService specs.TokenService,
	passService specs.PasswordService,
) *AuthRouter {
	return &AuthRouter{
		mux:          mux,
		userStore:    userStore,
		tokenService: tokenService,
		passService:  passService,
	}
}

func (r *AuthRouter) RegisterRoutes() {
	r.mux.HandleFunc("POST /auth/signup", r.signup)
	r.mux.HandleFunc("POST /auth/authenticate", r.authenticate)
	r.mux.HandleFunc("POST /auth/token/refresh", r.refreshTokens)
}

// create a new user, and return the created user information.
// status code: 201 Created
func (router *AuthRouter) signup(w http.ResponseWriter, r *http.Request) {
	cmd := &commands.SignupCmd{}

	if err := json.NewDecoder(r.Body).Decode(&cmd); err != nil {
		response.WriteError(w, err)
		return
	}

	if err := cmd.Validate(); err != nil {
		response.WriteError(w, err)
		return
	}

	user, err := handlers.SignupHandler(r.Context(), cmd, router.userStore, router.passService)
	if err != nil {
		response.WriteError(w, err)
		return
	}

	response.WriteJSON(w, http.StatusCreated, user)
}

func (router *AuthRouter) authenticate(w http.ResponseWriter, r *http.Request) {
	cmd := &commands.AuthCreditianls{}

	if err := json.NewDecoder(r.Body).Decode(&cmd); err != nil {
		response.WriteError(w, err)
		return
	}

	if err := cmd.Validate(); err != nil {
		response.WriteError(w, err)
		return
	}

	tokenPair, err := handlers.Authenticate(r.Context(), cmd, router.userStore, router.tokenService, router.passService)
	if err != nil {
		response.WriteError(w, err)
		return
	}

	response.WriteJSON(w, http.StatusOK, tokenPair)
}

func (router *AuthRouter) refreshTokens(w http.ResponseWriter, r *http.Request) {

	cmd := &commands.RefreshTokenCommand{}

	if err := json.NewDecoder(r.Body).Decode(&cmd); err != nil {
		response.WriteError(w, err)
		return
	}

	if err := cmd.Validate(); err != nil {
		response.WriteError(w, err)
		return
	}

	tokenPair, err := handlers.RefreshTokenPair(r.Context(), cmd, router.tokenService)
	if err != nil {
		response.WriteError(w, err)
		return
	}

	response.WriteJSON(w, http.StatusOK, tokenPair)
}
