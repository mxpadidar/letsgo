package response

import (
	"net/http"

	"github.com/mxpadidar/letsgo/internal/domain/errors"
)

var errStatusCode map[errors.ErrType]int = map[errors.ErrType]int{
	errors.ErrValidation: http.StatusBadRequest,
	errors.ErrNotFound:   http.StatusNotFound,
	errors.ErrAuthFailed: http.StatusUnauthorized,
	errors.ErrPermDenied: http.StatusForbidden,
	errors.ErrInternal:   http.StatusInternalServerError,
}

func getErrStatusCode(err *errors.Err) int {
	code, ok := errStatusCode[err.Typ]
	if !ok {
		return http.StatusInternalServerError
	}
	return code
}
