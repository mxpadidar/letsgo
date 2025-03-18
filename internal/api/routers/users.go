package routers

import (
	"net/http"

	"github.com/mxpadidar/letsgo/internal/api/response"
	"github.com/mxpadidar/letsgo/internal/core/dtos"
	"github.com/mxpadidar/letsgo/internal/core/specs"
	"github.com/mxpadidar/letsgo/internal/core/stores"
	"github.com/mxpadidar/letsgo/internal/core/types"
)

type UsersRouter struct {
	mux         *http.ServeMux
	userStore   stores.UserStore
	passService specs.PasswordService
}

func NewUsersRouter(
	mux *http.ServeMux,
	userStore stores.UserStore,
	passService specs.PasswordService,
) *UsersRouter {
	return &UsersRouter{
		mux:         mux,
		userStore:   userStore,
		passService: passService,
	}
}

func (router *UsersRouter) RegisterRoutes() {
	router.mux.HandleFunc("/users/me", router.getMe)
}

func (router *UsersRouter) getMe(w http.ResponseWriter, r *http.Request) {
	userDto, ok := r.Context().Value(types.UserContextKey).(*dtos.UserDto)
	if !ok {
		response.WriteError(w, types.ErrResourceNotFound)
		return
	}
	response.WriteJSON(w, http.StatusOK, userDto)
}
