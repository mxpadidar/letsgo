package server

import (
	"net/http"

	"github.com/mxpadidar/letsgo/internal/api/handlers"
	"github.com/mxpadidar/letsgo/internal/api/types"
	"github.com/mxpadidar/letsgo/internal/core/services"
)

type Server struct {
	logger      services.LogService
	authHandler *handlers.AuthHandler
	userHandler *handlers.UserHandler
	middlewares []types.Middleware
	authz       types.AuthzMiddleware
}

func NewServer(logger services.LogService, authHandler *handlers.AuthHandler, userHandler *handlers.UserHandler, authz types.AuthzMiddleware, middlewares ...types.Middleware) *Server {
	return &Server{
		logger:      logger,
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
	s.logger.Infof("Server started on port 8000")
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
