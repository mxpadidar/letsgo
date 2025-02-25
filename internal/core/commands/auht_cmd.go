package commands

import (
	"context"
	"fmt"

	"github.com/mxpadidar/letsgo/internal/core/dtos"
	"github.com/mxpadidar/letsgo/internal/core/services"
	"github.com/mxpadidar/letsgo/internal/core/stores"
	"github.com/mxpadidar/letsgo/internal/core/types"
)

type AuthCommand struct {
	userStore stores.UserStore
	hashSrvc  services.HashService
	Username  string `json:"username"`
	Password  string `json:"password"`
}

func NewAuthCmd(userStore stores.UserStore, hashSrvc services.HashService) *AuthCommand {
	return &AuthCommand{
		userStore: userStore,
		hashSrvc:  hashSrvc,
	}
}

func (cmd *AuthCommand) Execute(ctx context.Context) (*dtos.UserDto, error) {
	user, err := cmd.userStore.FindByUsername(ctx, cmd.Username)
	if err != nil && err != types.ErrNotFound {
		println(err)
		return nil, err
	}

	if user == nil {
		println("not found")
		return nil, types.ErrNotFound
	}

	println(cmd.Password)

	cmdPassHash, err := cmd.hashSrvc.HashPassword(cmd.Password)
	if err != nil {
		println(err)
		return nil, err
	}

	fmt.Printf("cmdPassHash: %s\n", cmdPassHash)
	fmt.Printf("userPassHash: %s\n", user.HashPassword)

	isMatch := cmd.hashSrvc.ComparePassword(user.HashPassword, cmd.Password)

	println(isMatch)
	if !isMatch {
		return nil, types.ErrInvalidCredentials
	}
	return dtos.NewUserDtoFromUser(user), nil
}
