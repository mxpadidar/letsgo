package helpers

import (
	"fmt"
	"net/http"

	"github.com/mxpadidar/letsgo/internal/domain/dtos"
	"github.com/mxpadidar/letsgo/internal/domain/errors"
)

func GetRequestPaginate(r *http.Request) (*dtos.PaginateDto, error) {
	query := r.URL.Query()
	params := make(map[string]string)

	for key, values := range query {
		if len(values) > 1 {
			return nil, errors.NewErr(errors.ErrValidation, fmt.Sprintf("multiple values for parameter %s", key), nil)
		}
		params[key] = values[0]
	}

	return dtos.NewPaginateDtoFromMap(params)
}
