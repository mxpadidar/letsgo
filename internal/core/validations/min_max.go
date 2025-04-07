package validations

import (
	"github.com/mxpadidar/letsgo/internal/core/errors"
)

func MinMaxValidation(key, val string, min, max int) error {
	if len(val) < min || len(val) > max {
		return errors.NewValidationErr("`%s` must be between %d and %d characters long", key, min, max)
	}
	return nil
}
