package types

import (
	"net/http"

	"github.com/mxpadidar/letsgo/internal/core/types"
)

type Middleware func(http.Handler) http.Handler

type AuthzMiddleware func(types.Permission, http.HandlerFunc) http.HandlerFunc
