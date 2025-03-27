package commands

import (
	"context"
	"strings"

	"github.com/mxpadidar/letsgo/internal/domain/entities"
	"github.com/mxpadidar/letsgo/internal/domain/enums"
	"github.com/mxpadidar/letsgo/internal/domain/errors"
	"github.com/mxpadidar/letsgo/internal/domain/stores"
	"github.com/mxpadidar/letsgo/internal/domain/validations"
)

type SignupCmd struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (cmd *SignupCmd) Execute(ctx context.Context, userStore stores.UserStore) (*entities.User, *errors.Err) {
	if err := cmd.validate(); err != nil {
		return nil, err
	}

	if user, _ := userStore.GetByUsername(ctx, cmd.Username); user != nil {
		return nil, errors.NewErr(errors.ErrConflict, "username already exists", nil)
	}

	user := entities.NewUser(cmd.Username, cmd.Password, enums.RoleUser)
	if err := userStore.Persist(ctx, user); err != nil {
		return nil, err
	}
	return user, nil
}

func (cmd *SignupCmd) validate() *errors.Err {
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
