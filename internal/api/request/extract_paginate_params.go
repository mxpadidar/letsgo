package request

import (
	"net/http"

	"github.com/mxpadidar/letsgo/internal/core/errors"
	"github.com/mxpadidar/letsgo/internal/core/types"
)

// ExtractPaginateParams extracts pagination parameters from an HTTP request.
// Returns a Paginate object and an error if any.
func ExtractPaginateParams(r *http.Request) (*types.Paginate, error) {
	query := r.URL.Query()
	params := make(map[string]string)

	for key, values := range query {
		if len(values) > 1 {
			return nil, errors.NewValidationErr("multiple values for parameter %s", key)
		}
		params[key] = values[0]
	}

	return types.NewPaginateFromMap(params)
}
