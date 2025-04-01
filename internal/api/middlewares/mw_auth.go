package middlewares

import (
	"context"
	"net/http"
	"slices"

	"github.com/mxpadidar/letsgo/internal/api/request"
	"github.com/mxpadidar/letsgo/internal/api/response"
	"github.com/mxpadidar/letsgo/internal/domain/services"
	"github.com/mxpadidar/letsgo/internal/domain/types"
)

func AuthMiddleware(tokenService services.TokenService) Middleware {
	allowedPaths := []string{"/auth/login", "/auth/signup"}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Allow public routes
			if slices.Contains(allowedPaths, r.URL.Path) {
				next.ServeHTTP(w, r)
				return
			}

			// Extract token
			token, err := request.ExtractBearerToken(r)
			if err != nil {
				response.WriteError(w, err)
				return
			}

			// Decode token
			authUser, err := tokenService.Decode(r.Context(), token)
			if err != nil {
				response.WriteError(w, err)
				return
			}

			// Store user in context
			ctx := context.WithValue(r.Context(), types.AuthUserKey, authUser)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
