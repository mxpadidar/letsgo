package bootstrap

import (
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/mxpadidar/letsgo/internal/api/handlers"
	"github.com/mxpadidar/letsgo/internal/api/middlewares"
	"github.com/mxpadidar/letsgo/internal/api/server"
	"github.com/mxpadidar/letsgo/internal/domain/services"
	"github.com/mxpadidar/letsgo/internal/domain/types"
	infraServices "github.com/mxpadidar/letsgo/internal/infra/services"
	"github.com/mxpadidar/letsgo/internal/infra/stores"
)

// Bootstrap initializes and returns the server
func Bootstrap() *server.Server {
	// Load Config
	configs, err := LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	// Connect to Database
	db := sqlx.MustConnect("postgres", configs.PostgresDSN)
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	// Initialize Dependencies
	userStore := stores.NewUserDBStore(db)
	hashService := infraServices.NewBcryptHash()
	jwtService := infraServices.NewJwtService(configs.JWTSecret, configs.AccessTokenTTL)
	stdLogService := infraServices.NewStdLogService()

	authService := services.NewAuthService(userStore, hashService, jwtService)
	userService := services.NewUserService(userStore)

	permService := services.NewPermService(
		services.RolePermsMap{
			types.RoleAdmin:  types.PermUserAll,
			types.RoleMember: types.PermUserRead,
			types.RoleGuest:  types.PermUserRead,
		},
		stdLogService,
	)

	authHandler := handlers.NewAuthHandler(authService)
	userHandler := handlers.NewUserHandler(userService)

	// Initialize Middlewares
	logMw := middlewares.LogMiddlewareFactory(stdLogService)
	authMw := middlewares.AuthMiddlewareFactory(jwtService)
	authzMw := middlewares.AuthzMiddlewareFactory(permService)

	// Create Server
	return server.NewServer(authHandler, userHandler, authzMw, logMw, authMw)
}
