package services

import (
	"context"

	"github.com/mxpadidar/letsgo/internal/domain/entities"
	"github.com/mxpadidar/letsgo/internal/domain/types"
)

type TokenService interface {
	Encode(ctx context.Context, user *entities.User) (token *types.Token, err error)
	Decode(ctx context.Context, tokenString string) (user *types.AuthUser, err error)
}
