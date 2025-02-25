package server

import (
	"net/http"

	"github.com/mxpadidar/letsgo/internal/api/handlers"
	"github.com/mxpadidar/letsgo/internal/core/services"
	"github.com/mxpadidar/letsgo/internal/core/stores"
)

type Server struct {
	addr        string
	router      *http.ServeMux
	userStore   stores.UserStore
	hashService services.HashService
}

func NewServer(addr string, userStore stores.UserStore, hashService services.HashService) *Server {
	router := http.NewServeMux()
	return &Server{
		addr:        addr,
		userStore:   userStore,
		hashService: hashService,
		router:      router,
	}

}

func (s *Server) SetupHandlers() {
	s.router.Handle("/health", http.HandlerFunc(s.healthCheck))

	auth := handlers.NewAuthHandler(
		s.router,
		s.userStore,
		s.hashService,
	)

	auth.SetupRoutes()
}

func (s *Server) Start() error {
	return http.ListenAndServe(s.addr, s.router)
}

func (s *Server) healthCheck(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
}
