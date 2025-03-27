package validations

import (
	"unicode"

	"github.com/mxpadidar/letsgo/internal/domain/errors"
)

func UsernameValidation(username string) *errors.Err {

	if len(username) == 0 {
		return errors.NewErr(errors.ErrValidation, "username is required", nil)
	}

	if err := MinMaxValidation("username", username, 6, 20); err != nil {
		return err
	}

	if !unicode.IsLetter(rune(username[0])) {
		return errors.NewErr(errors.ErrValidation, "username must start with a letter", nil)
	}

	for _, r := range username {
		if !unicode.IsLetter(r) && !unicode.IsNumber(r) && r != '_' {
			return errors.NewErr(errors.ErrValidation, "username can only contain letters, numbers, and underscore", nil)
		}
	}

	return nil
}
