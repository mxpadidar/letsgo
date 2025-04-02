package types

import (
	"net/http"

	"github.com/mxpadidar/letsgo/internal/domain/types"
)

type Middleware func(http.Handler) http.Handler

type Authz func(types.Permission, http.HandlerFunc) http.HandlerFunc
