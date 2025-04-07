package bootstrap

import (
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/mxpadidar/letsgo/internal/api/handlers"
	"github.com/mxpadidar/letsgo/internal/api/middlewares"
	"github.com/mxpadidar/letsgo/internal/api/server"
	"github.com/mxpadidar/letsgo/internal/core/services"
	"github.com/mxpadidar/letsgo/internal/core/types"
	"github.com/mxpadidar/letsgo/internal/infrastructure/adapters"
)

// Bootstrap initializes and returns the server
func Bootstrap() *server.Server {
	// Load Config
	logger := adapters.NewSlogLogger()
	configs, err := LoadConfig(logger)
	if err != nil {
		log.Fatal(err)
	}

	// Connect to Database
	db := sqlx.MustConnect("postgres", configs.PostgresDSN)
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	// Initialize Dependencies
	store := NewStore(db, logger)

	hashService := adapters.NewBcryptHash(logger)
	jwtService := adapters.NewJwtService(configs.AccessTokenSecret, configs.RefreshTokenSecret, configs.AccessTokenTTL, configs.RefreshTokenTTL)

	authService := services.NewAuthService(logger, store.Users, store.Permits, hashService, jwtService)
	userService := services.NewUserService(logger, store.Users)

	permService := services.NewPermService(
		services.RolePermsMap{
			types.RoleAdmin:  types.PermUserAll,
			types.RoleMember: types.PermUserRead,
			types.RoleGuest:  types.PermUserRead,
		},
		logger,
	)

	authHandler := handlers.NewAuthHandler(authService)
	userHandler := handlers.NewUserHandler(userService)

	// Initialize Middlewares
	logMw := middlewares.LogMiddlewareFactory(logger)
	authMw := middlewares.AuthMiddlewareFactory(jwtService, logger)
	authzMw := middlewares.AuthzMiddlewareFactory(permService, logger)

	// Create Server
	return server.NewServer(logger, authHandler, userHandler, authzMw, logMw, authMw)
}
