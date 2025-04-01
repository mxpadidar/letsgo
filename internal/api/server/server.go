package server

import (
	"log"
	"net/http"

	"github.com/mxpadidar/letsgo/internal/api/handlers"
	"github.com/mxpadidar/letsgo/internal/api/middlewares"
)

type Server struct {
	authHandler *handlers.AuthHandler
	userHandler *handlers.UserHandler
	middlewares middlewares.Middleware
}

func NewServer(authHandler *handlers.AuthHandler, userHandler *handlers.UserHandler, authMiddleware middlewares.Middleware) *Server {
	return &Server{authHandler: authHandler, userHandler: userHandler, middlewares: authMiddleware}
}

func (s *Server) Start() error {
	mux := http.NewServeMux()
	handler := s.middlewares(mux)

	s.authHandler.RegisterRoutes(mux)
	s.userHandler.RegisterRoutes(mux)
	log.Printf("Server started on port %s", ":8000")
	return http.ListenAndServe(":8000", handler)
}
