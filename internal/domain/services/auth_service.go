package services

import (
	"context"
	"fmt"

	"github.com/mxpadidar/letsgo/internal/domain/commands"
	"github.com/mxpadidar/letsgo/internal/domain/entities"
	"github.com/mxpadidar/letsgo/internal/domain/errors"
	"github.com/mxpadidar/letsgo/internal/domain/stores"
	"github.com/mxpadidar/letsgo/internal/domain/types"
)

type AuthService struct {
	userStore      stores.UserStore
	passwordHasher HashService
	tokenService   TokenService
}

func NewAuthService(userStore stores.UserStore, passwordHasher HashService, tokenService TokenService) *AuthService {
	return &AuthService{userStore: userStore, passwordHasher: passwordHasher, tokenService: tokenService}
}

func (h *AuthService) Login(ctx context.Context, cmd *commands.LoginCommand) (*types.Token, error) {
	user, err := h.userStore.GetByUsername(ctx, cmd.Username)
	if err != nil {
		return nil, err
	}

	if user == nil {
		errMsg := fmt.Sprintf("user with username %s not found", cmd.Username)
		return nil, errors.NewNotFoundError(errMsg)
	}

	if err := h.passwordHasher.Compare(ctx, user.HashedPassword, cmd.Password); err != nil {
		return nil, err
	}
	return h.tokenService.Encode(ctx, user)

}

func (h *AuthService) Signup(ctx context.Context, cmd *commands.SignupCommand) (*entities.User, error) {
	if err := cmd.Validate(); err != nil {
		return nil, err
	}

	if user, _ := h.userStore.GetByUsername(ctx, cmd.Username); user != nil {
		errMsg := fmt.Sprintf("user with username %s already exists", cmd.Username)
		return nil, errors.NewConflictError(errMsg)
	}

	hashedPassword, err := h.passwordHasher.Hash(ctx, cmd.Password)
	if err != nil {
		return nil, err
	}

	user := entities.NewUser(cmd.Username, hashedPassword, types.RoleMember)
	if err := h.userStore.Persist(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}
