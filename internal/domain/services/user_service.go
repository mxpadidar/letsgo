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
	authUser, ok := ctx.Value(types.AuthUserKey).(*types.AuthUser)
	if !ok {
		return nil, errors.NewAuthFailedError("authentication required")
	}

	user, err := h.userStore.GetByUsername(ctx, authUser.Username)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (h *UserService) ListUsers(ctx context.Context, paginate *types.Paginate) ([]*entities.User, error) {
	return h.userStore.List(ctx, paginate)
}
