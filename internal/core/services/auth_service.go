package services

import (
	"context"

	"github.com/mxpadidar/letsgo/internal/core/commands"
	"github.com/mxpadidar/letsgo/internal/core/entities"
	"github.com/mxpadidar/letsgo/internal/core/errors"
	"github.com/mxpadidar/letsgo/internal/core/stores"
	"github.com/mxpadidar/letsgo/internal/core/types"
)

type AuthService struct {
	Logger       LogService
	userStore    stores.UserStore
	permitStore  stores.PermitStore
	hashService  HashService
	tokenService TokenService
}

func NewAuthService(logger LogService, userStore stores.UserStore, permitStore stores.PermitStore, passwordHasher HashService, tokenService TokenService) *AuthService {
	return &AuthService{Logger: logger, userStore: userStore, permitStore: permitStore, hashService: passwordHasher, tokenService: tokenService}
}

func (s *AuthService) Signup(ctx context.Context, cmd *commands.SignupCommand) (*entities.User, error) {
	if err := cmd.Validate(); err != nil {
		return nil, err
	}

	if user, _ := s.userStore.GetByUsername(ctx, cmd.Username); user != nil {
		return nil, errors.NewConflictErr("user with username `%s` already exists", cmd.Username)
	}

	hashedPassword, err := s.hashService.Hash(ctx, cmd.Password)
	if err != nil {
		return nil, err
	}

	user := entities.NewUser(cmd.Username, hashedPassword, types.RoleMember)
	if err := s.userStore.Persist(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *AuthService) CreatePermit(ctx context.Context, cmd *commands.CreatePermitCmd) (*types.TokenPair, error) {
	user, err := s.userStore.GetByUsername(ctx, cmd.Username)
	if err != nil || user == nil {
		return nil, errors.AuthErr
	}

	if err := s.hashService.Compare(ctx, user.HashedPassword, cmd.Password); err != nil {
		return nil, err
	}
	permit, err := s.permitStore.Create(ctx, user.ID, user.Role)
	if err != nil {
		return nil, err
	}
	return s.tokenService.GenerateTokenPair(ctx, permit)
}

func (s *AuthService) RotatePermit(ctx context.Context, cmd *commands.RotatePermitCmd) (*types.TokenPair, error) {
	permit, err := s.tokenService.DecodeRefreshToken(ctx, cmd.RefreshToken)
	if err != nil {
		s.Logger.Errorf("failed to decode refresh token: %v", err)
		return nil, err
	}

	newPermit, err := s.permitStore.Rotate(ctx, permit.ID)
	if err != nil {
		if errors.GetErrorType(err) == errors.ErrNotFound {
			return nil, errors.AuthErr
		}
		return nil, err
	}
	return s.tokenService.GenerateTokenPair(ctx, newPermit)
}

func (s *AuthService) RevokePermit(ctx context.Context) error {
	permit, ok := ctx.Value(types.PermitContextKey).(*entities.Permit)
	if !ok {
		return errors.AuthErr
	}
	if err := s.permitStore.Delete(ctx, permit.ID); err != nil {
		return err
	}
	return nil
}
