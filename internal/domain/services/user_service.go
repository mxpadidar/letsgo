package services

import (
	"context"

	"github.com/mxpadidar/letsgo/internal/domain/entities"
	"github.com/mxpadidar/letsgo/internal/domain/errors"
	"github.com/mxpadidar/letsgo/internal/domain/stores"
	"github.com/mxpadidar/letsgo/internal/domain/types"
)

type UserService struct {
	userStore stores.UserStore
}

func NewUserService(userStore stores.UserStore) *UserService {
	return &UserService{userStore: userStore}
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
