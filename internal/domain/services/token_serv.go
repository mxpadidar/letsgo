package services

import (
	"context"

	"github.com/mxpadidar/letsgo/internal/domain/dtos"
	"github.com/mxpadidar/letsgo/internal/domain/entities"
)

type TokenService interface {
	Encode(ctx context.Context, user *entities.User) (token *dtos.TokenDTO, err error)
	Decode(ctx context.Context, tokenString string) (payload *dtos.TokenPayload, err error)
}
