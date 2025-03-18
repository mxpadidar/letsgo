package commands

import "github.com/mxpadidar/letsgo/internal/core/types"

type AuthCreditianls struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (cmd *AuthCreditianls) Validate() error {
	if cmd.Username == "" {
		return types.ErrValidation
	}

	if cmd.Password == "" {
		return types.ErrValidation
	}

	return nil
}
