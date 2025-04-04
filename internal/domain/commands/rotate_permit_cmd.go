package commands

import "github.com/mxpadidar/letsgo/internal/domain/errors"

type RotatePermitCmd struct {
	RefreshToken string `json:"refresh_token"`
}

func (cmd *RotatePermitCmd) Validate() error {
	if cmd.RefreshToken == "" {
		return errors.NewValidationErr("refresh_token is required")
	}
	return nil
}
