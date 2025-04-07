package response

import (
	"encoding/json"
	"net/http"

	"github.com/mxpadidar/letsgo/internal/core/errors"
	"github.com/mxpadidar/letsgo/internal/core/services"
)

func WriteError(w http.ResponseWriter, logger services.LogService, err error) {
	appErr, ok := err.(*errors.AppError)
	if !ok {
		logger.Errorf("failed to convert error to AppError: %v", err)
		appErr = errors.InternalErr
	}

	status := getErrStatusCode(appErr)
	data := map[string]string{"message": appErr.Message}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		logger.Errorf("failed to write error response: %v", err)

	}
}
