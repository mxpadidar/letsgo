package server

import (
	"net/http"

	"github.com/mxpadidar/letsgo/internal/api/middlewares"
	"github.com/mxpadidar/letsgo/internal/api/routers"
	"github.com/mxpadidar/letsgo/internal/core/specs"
	"github.com/mxpadidar/letsgo/internal/core/stores"
)

type Server struct {
	mux             *http.ServeMux
	addr            string
	tokenService    specs.TokenService
	passwordService specs.PasswordService
	userStore       stores.UserStore
	mw              middlewares.Middleware
}

func NewServer(addr string, tokenService specs.TokenService, passwordService specs.PasswordService, userStore stores.UserStore, mw middlewares.Middleware) *Server {
	return &Server{
		mux:             http.NewServeMux(),
		addr:            addr,
		tokenService:    tokenService,
		passwordService: passwordService,
		userStore:       userStore,
		mw:              mw,
	}
}

func (s *Server) Setup() {
	auth := routers.NewAuthRouter(s.mux, s.userStore, s.tokenService, s.passwordService)
	auth.RegisterRoutes()
	users := routers.NewUsersRouter(s.mux, s.userStore, s.passwordService)
	users.RegisterRoutes()
}

func (s *Server) Start() error {
	return http.ListenAndServe(s.addr, s.mw(s.mux))
}
