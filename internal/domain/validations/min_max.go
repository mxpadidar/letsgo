package validations

import (
	"fmt"

	"github.com/mxpadidar/letsgo/internal/domain/errors"
)

func MinMaxValidation(key, val string, min, max int) *errors.Err {
	if len(val) < min || len(val) > max {
		msg := fmt.Sprintf("%s must be between %d and %d characters long", key, min, max)
		return errors.NewErr(errors.ErrValidation, msg, nil)
	}
	return nil
}
