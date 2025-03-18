package handlers

import (
	"context"
	"strconv"

	"github.com/mxpadidar/letsgo/internal/core/commands"
	"github.com/mxpadidar/letsgo/internal/core/dtos"
	"github.com/mxpadidar/letsgo/internal/core/specs"
	"github.com/mxpadidar/letsgo/internal/core/stores"
)

func Authenticate(ctx context.Context, cmd *commands.AuthCreditianls, userStore stores.UserStore, tokenSrv specs.TokenService, passService specs.PasswordService) (*dtos.TokenPair, error) {
	user, err := userStore.FindByUsername(ctx, cmd.Username)
	if err != nil {
		return nil, err
	}
	if err := passService.Verify(user.HashPassword, cmd.Password); err != nil {
		return nil, err
	}
	sub := strconv.Itoa(user.ID)
	return tokenSrv.GenerateTokenPair(sub)
}
