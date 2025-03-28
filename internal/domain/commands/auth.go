package commands

import (
	"context"
	"strings"

	"github.com/mxpadidar/letsgo/internal/domain/entities"
	"github.com/mxpadidar/letsgo/internal/domain/errors"
	"github.com/mxpadidar/letsgo/internal/domain/services"
	"github.com/mxpadidar/letsgo/internal/domain/stores"
	"github.com/mxpadidar/letsgo/internal/domain/validations"
)

type AuthCmd struct {
	Username string
	Password string
}

func (cmd *AuthCmd) Execute(
	ctx context.Context,
	userStore stores.UserStore,
	passwordServ services.PasswordService,
) (*entities.User, error) {
	if err := cmd.validate(); err != nil {
		return nil, err
	}

	user, err := userStore.GetByUsername(ctx, cmd.Username)
	if err != nil {
		return nil, err
	}

	if err := passwordServ.Compare(ctx, user.HashedPassword, cmd.Password); err != nil {
		return nil, err
	}

	return user, nil
}

func (cmd *AuthCmd) validate() *errors.Err {
	// update username to lowercase and strip whitespace
	cmd.Username = strings.ToLower(cmd.Username)
	cmd.Username = strings.TrimSpace(cmd.Username)

	if err := validations.UsernameValidation(cmd.Username); err != nil {
		return err
	}

	if cmd.Password == "" {
		return errors.NewErr(errors.ErrValidation, "password is required", nil)
	}

	if err := validations.MinMaxValidation("password", cmd.Password, 6, 20); err != nil {
		return err
	}

	return nil
}
