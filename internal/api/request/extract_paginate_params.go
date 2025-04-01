package request

import (
	"fmt"
	"net/http"

	"github.com/mxpadidar/letsgo/internal/domain/errors"
	"github.com/mxpadidar/letsgo/internal/domain/types"
)

// ExtractPaginateParams extracts pagination parameters from an HTTP request.
// Returns a Paginate object and an error if any.
func ExtractPaginateParams(r *http.Request) (*types.Paginate, error) {
	query := r.URL.Query()
	params := make(map[string]string)

	for key, values := range query {
		if len(values) > 1 {
			errMsg := fmt.Sprintf("multiple values for parameter %s", key)
			return nil, errors.NewValidationError(errMsg)
		}
		params[key] = values[0]
	}

	return types.NewPaginateFromMap(params)
}
