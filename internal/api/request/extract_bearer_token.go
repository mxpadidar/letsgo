package request

import (
	"net/http"
	"strings"

	"github.com/mxpadidar/letsgo/internal/core/errors"
)

// ExtractBearerToken extracts the Bearer token from the Authorization header.
// Returns an error if the header is missing or malformed.
func ExtractBearerToken(r *http.Request) (string, error) {
	bearerToken := r.Header.Get("Authorization")

	if bearerToken == "" {
		return "", errors.AuthErr
	}

	parts := strings.Split(bearerToken, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return "", errors.AuthErr
	}

	return parts[1], nil
}
