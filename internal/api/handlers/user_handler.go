package handlers

import (
	"net/http"

	"github.com/mxpadidar/letsgo/internal/api/request"
	"github.com/mxpadidar/letsgo/internal/api/response"
	"github.com/mxpadidar/letsgo/internal/domain/services"
)

type UserHandler struct {
	users *services.UserService
}

func NewUserHandler(handler *services.UserService) *UserHandler {
	return &UserHandler{users: handler}
}

func (h *UserHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /users/me", h.getCurrentUser)
	mux.HandleFunc("/users", h.listUsers)
}

func (h *UserHandler) getCurrentUser(w http.ResponseWriter, r *http.Request) {
	user, err := h.users.GetCurrentUser(r.Context())
	if err != nil {
		response.WriteError(w, err)
		return
	}

	response.WriteJSON(w, http.StatusOK, user)
}

func (h *UserHandler) listUsers(w http.ResponseWriter, r *http.Request) {
	paginate, err := request.ExtractPaginateParams(r)
	if err != nil {
		response.WriteError(w, err)
		return
	}

	users, err := h.users.ListUsers(r.Context(), paginate)
	if err != nil {
		response.WriteError(w, err)
		return
	}

	response.WriteJSON(w, http.StatusOK, users)
}
