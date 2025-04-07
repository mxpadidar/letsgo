package response

import (
	"net/http"

	"github.com/mxpadidar/letsgo/internal/core/errors"
)

func getErrStatusCode(err *errors.AppError) int {

	var errStatusCode map[errors.ErrorType]int = map[errors.ErrorType]int{
		errors.ErrValidation:   http.StatusBadRequest,
		errors.ErrNotFound:     http.StatusNotFound,
		errors.ErrAuthFailed:   http.StatusUnauthorized,
		errors.ErrAccessDenied: http.StatusForbidden,
		errors.ErrInternal:     http.StatusInternalServerError,
	}

	code, ok := errStatusCode[err.ErrType]
	if !ok {
		return http.StatusInternalServerError
	}

	return code
}
