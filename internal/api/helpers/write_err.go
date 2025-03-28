package helpers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/mxpadidar/letsgo/internal/domain/errors"
)

func WriteError(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")
	appErr, ok := err.(*errors.Err)
	if !ok {
		appErr = errors.NewErr(errors.ErrInternal, "internal server error", err)
	}

	w.WriteHeader(getErrStatusCode(appErr))

	data := map[string]interface{}{"message": appErr.Msg}

	if appErr.Err != nil && appErr.Typ != errors.ErrInternal {
		data["data"] = map[string]string{
			"error": appErr.Err.Error(),
		}
	}

	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("Failed to write error response: %v", err)
	}
}

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
