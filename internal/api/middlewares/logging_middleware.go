package middlewares

import (
	"net/http"

	"github.com/mxpadidar/letsgo/internal/core/services"
)

func LogMiddlewareFactory(logger services.LogService) func(next http.Handler) http.Handler {
	middleware := func(next http.Handler) http.Handler {
		handler := func(w http.ResponseWriter, r *http.Request) {
			logger.Infof("Request received: %s %s", r.Method, r.URL.Path)
			next.ServeHTTP(w, r)
		}

		return http.HandlerFunc(handler)
	}

	return middleware
}
