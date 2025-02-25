package commands

import (
	"context"

	"github.com/mxpadidar/letsgo/internal/core/dtos"
	"github.com/mxpadidar/letsgo/internal/core/entities"
	"github.com/mxpadidar/letsgo/internal/core/services"
	"github.com/mxpadidar/letsgo/internal/core/stores"
	"github.com/mxpadidar/letsgo/internal/core/types"
)

type RegisterCmd struct {
	Username    string `json:"username"`
	RawPassword string `json:"password"`
	Fname       string `json:"first_name"`
	Lname       string `json:"last_name"`
	userStore   stores.UserStore
	hashService services.HashService
}

func NewRegisterCmd(userStore stores.UserStore, hashService services.HashService) *RegisterCmd {
	return &RegisterCmd{
		userStore:   userStore,
		hashService: hashService,
	}
}

func (c *RegisterCmd) Execute(ctx context.Context) (*dtos.UserDto, error) {
	println(c.Username)
	exists, err := c.userStore.FindByUsername(ctx, c.Username)
	if err != nil && err != types.ErrNotFound {
		return nil, err
	}
	if exists != nil {
		return nil, types.ErrConflict
	}

	hashPassword, err := c.hashService.HashPassword(c.RawPassword)
	if err != nil {
		return nil, err
	}

	user := entities.NewUser(
		c.Username,
		hashPassword,
		c.Fname,
		c.Lname,
		false,
	)

	err = c.userStore.Save(ctx, user)
	if err != nil {
		return nil, err
	}

	return dtos.NewUserDtoFromUser(user), nil
}
