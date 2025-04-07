package middlewares

import (
	"net/http"

	"github.com/mxpadidar/letsgo/internal/api/response"
	"github.com/mxpadidar/letsgo/internal/core/entities"
	"github.com/mxpadidar/letsgo/internal/core/errors"
	"github.com/mxpadidar/letsgo/internal/core/services"
	"github.com/mxpadidar/letsgo/internal/core/types"
)

// AuthzMiddlewareFactory creates an authorization middleware
func AuthzMiddlewareFactory(permService *services.PermService, logger services.LogService) func(types.Permission, http.HandlerFunc) http.HandlerFunc {
	accessControl := func(perm types.Permission, next http.HandlerFunc) http.HandlerFunc {
		handler := func(w http.ResponseWriter, r *http.Request) {
			// Extract user from context
			permit, ok := r.Context().Value(types.PermitContextKey).(*entities.Permit)
			if !ok || permit == nil {
				response.WriteError(w, logger, errors.AuthErr)
				return
			}

			// Check if the user has the required permission
			if !permService.CheckPerm(permit.Role, perm) {
				response.WriteError(w, logger, errors.AccessErr)
				return
			}

			// Proceed to the next handler
			next.ServeHTTP(w, r)
		}
		return handler
	}
	return accessControl
}
