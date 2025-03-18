package handlers

import (
	"context"

	"github.com/mxpadidar/letsgo/internal/core/commands"
	"github.com/mxpadidar/letsgo/internal/core/dtos"
	"github.com/mxpadidar/letsgo/internal/core/entities"
	"github.com/mxpadidar/letsgo/internal/core/specs"
	"github.com/mxpadidar/letsgo/internal/core/stores"
	"github.com/mxpadidar/letsgo/internal/core/types"
)

func SignupHandler(ctx context.Context, cmd *commands.SignupCmd, userStore stores.UserStore, passwordService specs.PasswordService) (*dtos.UserDto, error) {
	exists, err := userStore.FindByUsername(ctx, cmd.Username)
	if err != nil && err != types.ErrResourceNotFound {
		return nil, err
	}
	if exists != nil {
		return nil, types.ErrConflict
	}

	hashPassword, err := passwordService.Hash(cmd.Password)
	if err != nil {
		return nil, err
	}

	user := entities.NewUser(
		cmd.Username,
		hashPassword,
		cmd.FirstName,
		cmd.LastName,
		false,
	)

	err = userStore.Save(ctx, user)
	if err != nil {
		return nil, err
	}

	return dtos.NewUserDtoFromUser(user), nil
}
