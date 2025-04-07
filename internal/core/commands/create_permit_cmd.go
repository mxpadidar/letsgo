package commands

import (
	"strings"

	"github.com/mxpadidar/letsgo/internal/core/errors"
	"github.com/mxpadidar/letsgo/internal/core/validations"
)

type CreatePermitCmd struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (cmd *CreatePermitCmd) Validate() error {
	// update username to lowercase and strip whitespace
	cmd.Username = strings.ToLower(cmd.Username)
	cmd.Username = strings.TrimSpace(cmd.Username)

	if err := validations.UsernameValidation(cmd.Username); err != nil {
		return err
	}

	if cmd.Password == "" {
		return errors.NewValidationErr("password is required")
	}

	if err := validations.MinMaxValidation("password", cmd.Password, 5, 20); err != nil {
		return err
	}

	return nil
}
