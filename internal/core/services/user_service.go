package services

import (
	"context"

	"github.com/mxpadidar/letsgo/internal/core/entities"
	"github.com/mxpadidar/letsgo/internal/core/errors"
	"github.com/mxpadidar/letsgo/internal/core/stores"
	"github.com/mxpadidar/letsgo/internal/core/types"
)

type UserService struct {
	Logger    LogService
	userStore stores.UserStore
}

func NewUserService(logger LogService, userStore stores.UserStore) *UserService {
	return &UserService{Logger: logger, userStore: userStore}
}

func (h *UserService) GetCurrentUser(ctx context.Context) (*entities.User, error) {
	permit, ok := ctx.Value(types.PermitContextKey).(*entities.Permit)
	if !ok {
		return nil, errors.AuthErr
	}

	user, err := h.userStore.GetByID(ctx, permit.UserID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (h *UserService) ListUsers(ctx context.Context, paginate *types.Paginate) ([]*entities.User, error) {
	return h.userStore.List(ctx, paginate)
}
