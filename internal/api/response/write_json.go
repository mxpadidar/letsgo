package response

import (
	"encoding/json"
	"log"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")

	if status != http.StatusNoContent && data == nil {
		log.Printf("status code %d requires non-nil data", status)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(status)

	if status == http.StatusNoContent {
		return
	}

	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("failed to encode response data: %v", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
}
