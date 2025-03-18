package response

import (
	"net/http"

	"github.com/mxpadidar/letsgo/internal/core/types"
)

var domainErrStatus = map[types.DomainError]int{
	types.ErrValidation:       http.StatusBadRequest,
	types.ErrResourceNotFound: http.StatusNotFound,
	types.ErrUnauthorized:     http.StatusUnauthorized,
	types.ErrForbidden:        http.StatusForbidden,
	types.ErrInternal:         http.StatusInternalServerError,
}

type errResponse struct {
	Error string `json:"error"`
}
