package request

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/mxpadidar/letsgo/internal/core/errors"
)

type validable interface {
	Validate() error
}

// ParseRequestBody parses the request body and validates it.
// It returns a context.Context and an error.
// cmd should be a pointer to a struct that implements Validatable interface.
func ParseRequestBody(r *http.Request, cmd validable) (context.Context, error) {
	if err := json.NewDecoder(r.Body).Decode(cmd); err != nil {
		return nil, errors.ValidationErr
	}

	if err := cmd.Validate(); err != nil {
		return nil, err
	}

	return r.Context(), nil
}
