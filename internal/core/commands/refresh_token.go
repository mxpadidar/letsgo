package commands

import (
	"github.com/mxpadidar/letsgo/internal/core/types"
)

type RefreshTokenCommand struct {
	RefreshToken string `json:"refresh_token"`
}

func (cmd *RefreshTokenCommand) Validate() error {
	if cmd.RefreshToken == "" {
		return types.ErrValidation
	}
	return nil
}
