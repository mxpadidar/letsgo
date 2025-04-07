package response

import (
	"encoding/json"
	"net/http"

	"github.com/mxpadidar/letsgo/internal/core/services"
)

func WriteJSON(w http.ResponseWriter, logger services.LogService, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if status == http.StatusNoContent || data == nil {
		return
	}

	if err := json.NewEncoder(w).Encode(data); err != nil {
		logger.Errorf("Failed to encode data: %v, %v", err, data)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "something went wrong; please report the issue"})
	}
}
