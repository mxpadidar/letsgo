package middlewares

import (
	"context"

	"net/http"
	"strconv"
	"strings"

	"github.com/mxpadidar/letsgo/internal/core/dtos"
	"github.com/mxpadidar/letsgo/internal/core/specs"
	"github.com/mxpadidar/letsgo/internal/core/stores"
	"github.com/mxpadidar/letsgo/internal/core/types"
)

func NewAuthMiddleware(ts specs.TokenService, store stores.UserStore) Middleware {

	return func(next http.Handler) http.Handler {

		handler := func(w http.ResponseWriter, r *http.Request) {
			if strings.HasPrefix(r.URL.Path, "/auth") {
				next.ServeHTTP(w, r)
				return
			}

			token := r.Header.Get("Authorization")

			if token == "" {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			parts := strings.Split(token, " ")
			if len(parts) != 2 || parts[0] != "Bearer" {
				http.Error(w, "Invalid token format", http.StatusUnauthorized)
				return
			}

			sub, err := ts.Decode(parts[1], "access")

			if err != nil {

				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			userId, err := strconv.Atoi(sub)
			if err != nil {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			ctx := r.Context()

			user, err := store.FindById(ctx, userId)
			if err != nil {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			userDto := dtos.NewUserDtoFromUser(user)
			ctx = context.WithValue(ctx, types.UserContextKey, userDto)

			next.ServeHTTP(w, r.WithContext(ctx))
		}

		return http.HandlerFunc(handler)
	}
}
