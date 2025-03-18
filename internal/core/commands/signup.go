package commands

import "github.com/mxpadidar/letsgo/internal/core/types"

type SignupCmd struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func (cmd *SignupCmd) Validate() error {
	if cmd.Username == "" {
		return types.ErrValidation
	}
	if cmd.Password == "" {
		return types.ErrValidation
	}
	return nil
}
