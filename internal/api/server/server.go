package server

import (
	"log"
	"net/http"

	"github.com/mxpadidar/letsgo/internal/api/routers"
	"github.com/mxpadidar/letsgo/internal/domain/services"
	"github.com/mxpadidar/letsgo/internal/domain/stores"
)

type Server struct {
	mux          *http.ServeMux
	UserStore    stores.UserStore
	TokenServ    services.TokenService
	PasswordServ services.PasswordService
}

func NewServer(userStore stores.UserStore, tokenServ services.TokenService, passwordServ services.PasswordService) *Server {
	return &Server{
		mux:          http.NewServeMux(),
		UserStore:    userStore,
		TokenServ:    tokenServ,
		PasswordServ: passwordServ,
	}
}

func (s *Server) Start(addr string) error {
	auth := routers.NewAuthRouter(s.mux, s.UserStore, s.TokenServ, s.PasswordServ)
	auth.Load()
	users := routers.NewUsersRouter(s.mux, s.UserStore)
	users.Load()
	log.Printf("Starting server on %s", addr)
	return http.ListenAndServe(addr, s.mux)
}
