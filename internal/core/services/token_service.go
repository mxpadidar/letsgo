package services

import (
	"context"

	"github.com/mxpadidar/letsgo/internal/core/entities"
	"github.com/mxpadidar/letsgo/internal/core/types"
)

type TokenService interface {
	GenerateTokenPair(ctx context.Context, permit *entities.Permit) (*types.TokenPair, error)
	DecodeRefreshToken(ctx context.Context, tokenString string) (*entities.Permit, error)
	DecodeAccessToken(ctx context.Context, tokenString string) (*entities.Permit, error)
}
