package response

import (
	"encoding/json"
	"log"
	"net/http"
)

func WriteOk(w http.ResponseWriter, msg string, resource interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	data := map[string]interface{}{"message": msg, "data": resource}
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("Failed to encode response: %v, %v", err, resource)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "something went wrong; please report the issue"})
	}
}
