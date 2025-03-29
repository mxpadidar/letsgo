package routers

import (
	"net/http"

	"github.com/mxpadidar/letsgo/internal/api/helpers"
	"github.com/mxpadidar/letsgo/internal/domain/entities"
	"github.com/mxpadidar/letsgo/internal/domain/errors"
	"github.com/mxpadidar/letsgo/internal/domain/queris"
	"github.com/mxpadidar/letsgo/internal/domain/stores"
	"github.com/mxpadidar/letsgo/internal/domain/types"
)

type UsersRouter struct {
	mux       *http.ServeMux
	userStore stores.UserStore
}

func NewUsersRouter(mux *http.ServeMux, userStore stores.UserStore) *UsersRouter {
	return &UsersRouter{mux: mux, userStore: userStore}
}

func (router *UsersRouter) Load() {
	router.mux.HandleFunc("GET /users", router.usersList)
	router.mux.HandleFunc("GET /users/me", router.getMe)
}

func (router *UsersRouter) usersList(w http.ResponseWriter, r *http.Request) {
	paginate, err := helpers.GetRequestPaginate(r)
	if err != nil {
		helpers.WriteError(w, err)
		return
	}

	query := queris.NewUsersListQuery(paginate)
	users, err := query.Fetch(r.Context(), router.userStore)
	if err != nil {
		helpers.WriteError(w, err)
		return
	}

	helpers.WriteJSON(w, users, http.StatusOK)
}

func (router *UsersRouter) getMe(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value(types.UserContextKey).(*entities.User)
	if !ok {
		helpers.WriteError(w, errors.NewErr(errors.ErrNotFound, "user not found", nil))
		return
	}

	helpers.WriteJSON(w, user, http.StatusOK)
}
