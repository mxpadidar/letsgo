package validations

import (
	"unicode"

	"github.com/mxpadidar/letsgo/internal/core/errors"
)

func UsernameValidation(username string) error {

	if len(username) == 0 {
		return errors.NewValidationErr("username cannot be empty")
	}

	if err := MinMaxValidation("username", username, 6, 20); err != nil {
		return err
	}

	if !unicode.IsLetter(rune(username[0])) {
		return errors.NewValidationErr("username must start with a letter")
	}

	for _, r := range username {
		if !unicode.IsLetter(r) && !unicode.IsNumber(r) && r != '_' {
			return errors.NewValidationErr("%c is not allowed in username", r)
		}
	}

	return nil
}
