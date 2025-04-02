package server

import (
	"log"
	"net/http"

	"github.com/mxpadidar/letsgo/internal/api/handlers"
	"github.com/mxpadidar/letsgo/internal/api/types"
)

type Server struct {
	authHandler *handlers.AuthHandler
	userHandler *handlers.UserHandler
	middlewares []types.Middleware
	authz       types.Authz
}

func NewServer(authHandler *handlers.AuthHandler, userHandler *handlers.UserHandler, authz types.Authz, middlewares ...types.Middleware) *Server {
	return &Server{
		authHandler: authHandler,
		userHandler: userHandler,
		middlewares: middlewares,
		authz:       authz,
	}
}

func (s *Server) Start() error {
	mux := http.NewServeMux()
	mwChain := s.ChainMiddlewares(s.middlewares...)

	s.authHandler.RegisterRoutes(mux)
	s.userHandler.RegisterRoutes(mux, s.authz)
	log.Printf("Server started on port %s", ":8000")
	return http.ListenAndServe(":8000", mwChain(mux))
}

func (s *Server) ChainMiddlewares(middlewares ...types.Middleware) types.Middleware {
	return func(next http.Handler) http.Handler {
		for i := len(middlewares) - 1; i >= 0; i-- {
			next = middlewares[i](next)
		}
		return next
	}
}
