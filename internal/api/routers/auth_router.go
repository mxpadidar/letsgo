package routers

import (
	"encoding/json"
	"net/http"
)

type AuthRouter struct {
	mux *http.ServeMux
}

func NewAuthRouter(mux *http.ServeMux) *AuthRouter {
	return &AuthRouter{
		mux: mux,
	}
}

func (ar *AuthRouter) Load() {
	ar.mux.HandleFunc("POST /auth/register", ar.registerUser)
}

func (ar *AuthRouter) registerUser(w http.ResponseWriter, r *http.Request) {
	// Set content type header first
	w.Header().Set("Content-Type", "application/json")

	// Set status code
	w.WriteHeader(http.StatusOK)

	// Write response
	if err := json.NewEncoder(w).Encode(map[string]string{"message": "User registered successfully"}); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
