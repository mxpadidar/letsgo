package middlewares

import (
	"net/http"

	"github.com/mxpadidar/letsgo/internal/api/response"
	"github.com/mxpadidar/letsgo/internal/domain/errors"
	"github.com/mxpadidar/letsgo/internal/domain/services"
	"github.com/mxpadidar/letsgo/internal/domain/types"
)

// AuthzMiddlewareFactory creates an authorization middleware
func AuthzMiddlewareFactory(permService *services.PermService) func(types.Permission, http.HandlerFunc) http.HandlerFunc {
	accessControl := func(perm types.Permission, next http.HandlerFunc) http.HandlerFunc {
		handler := func(w http.ResponseWriter, r *http.Request) {
			// Extract user from context
			authUser, ok := r.Context().Value(types.AuthUserKey).(*types.AuthUser)
			if !ok || authUser == nil {
				response.WriteError(w, errors.NewAuthFailedError("authentication failed"))
				return
			}

			// Check if the user has the required permission
			if !permService.CheckPerm(authUser.Role, perm) {
				response.WriteError(w, errors.NewAccessDeniedError("access denied"))
				return
			}

			// Proceed to the next handler
			next.ServeHTTP(w, r)
		}
		return handler
	}
	return accessControl
}
