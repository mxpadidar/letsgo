package response

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/mxpadidar/letsgo/internal/core/types"
)

func WriteError(w http.ResponseWriter, domainErr types.DomainError) {
	w.Header().Set("Content-Type", "application/json")
	code := domainErrStatus[domainErr]
	if code == 0 {
		code = http.StatusInternalServerError
		log.Printf("Error is not domainErr: %v", domainErr)
		domainErr = errors.New("Internal Error! Please try later again.")
	}
	w.WriteHeader(code)

	errResponse := errResponse{Error: domainErr.Error()}
	if err := json.NewEncoder(w).Encode(errResponse); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
