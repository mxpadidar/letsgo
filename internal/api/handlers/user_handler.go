package handlers

import (
	"net/http"

	"github.com/mxpadidar/letsgo/internal/api/request"
	"github.com/mxpadidar/letsgo/internal/api/response"
	apiTypes "github.com/mxpadidar/letsgo/internal/api/types"
	"github.com/mxpadidar/letsgo/internal/core/services"
	"github.com/mxpadidar/letsgo/internal/core/types"
)

type UserHandler struct {
	service *services.UserService
}

func NewUserHandler(service *services.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) RegisterRoutes(mux *http.ServeMux, authz apiTypes.AuthzMiddleware) {
	mux.HandleFunc("GET /users/me", authz(types.PermUserRead, h.getCurrentUser))
	mux.HandleFunc("GET /users", authz(types.PermUserAll, h.listUsers))
}

func (h *UserHandler) getCurrentUser(w http.ResponseWriter, r *http.Request) {
	user, err := h.service.GetCurrentUser(r.Context())
	if err != nil {
		response.WriteError(w, h.service.Logger, err)
		return
	}

	response.WriteJSON(w, h.service.Logger, http.StatusOK, user)
}

func (h *UserHandler) listUsers(w http.ResponseWriter, r *http.Request) {
	paginate, err := request.ExtractPaginateParams(r)
	if err != nil {
		response.WriteError(w, h.service.Logger, err)
		return
	}

	users, err := h.service.ListUsers(r.Context(), paginate)
	if err != nil {
		response.WriteError(w, h.service.Logger, err)
		return
	}

	response.WriteJSON(w, h.service.Logger, http.StatusOK, users)
}
