package middlewares

import (
	"context"

	"net/http"
	"strings"

	"github.com/mxpadidar/letsgo/internal/api/helpers"
	"github.com/mxpadidar/letsgo/internal/domain/errors"
	"github.com/mxpadidar/letsgo/internal/domain/services"
	"github.com/mxpadidar/letsgo/internal/domain/stores"
	"github.com/mxpadidar/letsgo/internal/domain/types"
)

func NewAuthErr(msg string) error {
	return errors.NewErr(errors.ErrAuthFailed, msg, nil)
}

func NewAuthMiddleware(tokenServ services.TokenService, userStore stores.UserStore) func(http.Handler) http.Handler {

	middleware := func(next http.Handler) http.Handler {

		handlerFunc := func(w http.ResponseWriter, r *http.Request) {
			if strings.HasPrefix(r.URL.Path, "/auth") {
				next.ServeHTTP(w, r)
				return
			}

			token := r.Header.Get("Authorization")

			if token == "" {
				helpers.WriteError(w, NewAuthErr("Authorization header is missing"))
				return
			}

			parts := strings.Split(token, " ")
			if len(parts) != 2 || parts[0] != "Bearer" {
				helpers.WriteError(w, NewAuthErr("Invalid token format"))
				return
			}

			payload, err := tokenServ.Decode(r.Context(), parts[1])
			if err != nil {
				helpers.WriteError(w, NewAuthErr("Auth failed"))
				return
			}

			ctx := r.Context()

			user, err := userStore.GetByID(ctx, payload.UserID)
			if err != nil {
				helpers.WriteError(w, NewAuthErr("Auth failed"))
				return
			}

			ctx = context.WithValue(ctx, types.UserContextKey, user)

			next.ServeHTTP(w, r.WithContext(ctx))
		}

		return http.HandlerFunc(handlerFunc)
	}

	return middleware
}
