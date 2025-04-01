package response

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/mxpadidar/letsgo/internal/domain/errors"
)

func WriteError(w http.ResponseWriter, err error) {
	appErr, ok := err.(*errors.AppError)
	if !ok {
		log.Printf("failed to convert error to AppError: %v", err)
		appErr = errors.NewInternalError("something goes wrong!")
	}

	status := getErrStatusCode(appErr)
	data := map[string]string{"message": appErr.Message}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("failed to write error response: %v", err)
	}
}
