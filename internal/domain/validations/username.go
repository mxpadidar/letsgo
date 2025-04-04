package validations

import (
	"fmt"
	"unicode"

	"github.com/mxpadidar/letsgo/internal/domain/errors"
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
			errMsg := fmt.Sprintf("%c is not allowed in username", r)
			return errors.NewValidationErr(errMsg)
		}
	}

	return nil
}
