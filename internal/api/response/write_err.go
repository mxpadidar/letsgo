package response

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/mxpadidar/letsgo/internal/domain/errors"
)

func WriteError(w http.ResponseWriter, err *errors.Err) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(getErrStatusCode(err))

	data := map[string]interface{}{"message": err.Msg}

	if err.Err != nil {
		data["data"] = map[string]string{
			"error": err.Err.Error(),
		}
	}

	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("Failed to write error response: %v", err)
	}
}
