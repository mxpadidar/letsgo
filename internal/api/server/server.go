package server

import (
	"log"
	"net/http"

	"github.com/mxpadidar/letsgo/internal/api/routers"
	"github.com/mxpadidar/letsgo/internal/domain/stores"
)

type Server struct {
	mux       *http.ServeMux
	UserStore stores.UserStore
}

func NewServer(userStore stores.UserStore) *Server {
	return &Server{
		mux:       http.NewServeMux(),
		UserStore: userStore,
	}
}

func (s *Server) Start(addr string) error {
	auth := routers.NewAuthRouter(s.mux)
	auth.Load()
	log.Printf("Starting server on %s", addr)
	return http.ListenAndServe(addr, s.mux)
}
