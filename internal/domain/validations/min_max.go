package validations

import (
	"fmt"

	"github.com/mxpadidar/letsgo/internal/domain/errors"
)

func MinMaxValidation(key, val string, min, max int) error {
	if len(val) < min || len(val) > max {
		errMsg := fmt.Sprintf("%s must be between %d and %d characters long", key, min, max)
		return errors.NewValidationError(errMsg)
	}
	return nil
}
