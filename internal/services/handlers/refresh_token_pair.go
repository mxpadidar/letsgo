package handlers

import (
	"context"

	"github.com/mxpadidar/letsgo/internal/core/commands"
	"github.com/mxpadidar/letsgo/internal/core/dtos"
	"github.com/mxpadidar/letsgo/internal/core/specs"
)

func RefreshTokenPair(ctx context.Context, cmd *commands.RefreshTokenCommand, tokenService specs.TokenService) (*dtos.TokenPair, error) {
	return tokenService.RefreshTokenPair(cmd.RefreshToken)
}
