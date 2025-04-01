package main

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/mxpadidar/letsgo/internal/api/handlers"
	"github.com/mxpadidar/letsgo/internal/api/middlewares"
	"github.com/mxpadidar/letsgo/internal/api/server"
	domainService "github.com/mxpadidar/letsgo/internal/domain/services"
	"github.com/mxpadidar/letsgo/internal/infra/configs"
	infraService "github.com/mxpadidar/letsgo/internal/infra/services"
	"github.com/mxpadidar/letsgo/internal/infra/stores"
)

func main() {
	env := infraService.NewEnvVarService()
	configs, err := configs.Load(env)
	if err != nil {
		log.Fatal(err)
	}

	db := sqlx.MustConnect("postgres", configs.PostgresDSN)
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	userStore := stores.NewUserDBStore(db)

	hash_service := infraService.NewBcryptHash()
	jwtService := infraService.NewJwtService(configs.JWTSecret, configs.AccessTokenTTL)

	authService := domainService.NewAuthService(userStore, hash_service, jwtService)
	userService := domainService.NewUserService(userStore)

	authHandler := handlers.NewAuthRouter(authService)
	userHandler := handlers.NewUserHandler(userService)

	middlewares := middlewares.ChainMiddlewares(middlewares.AuthMiddleware(jwtService), middlewares.LoggingMiddleware)

	server := server.NewServer(authHandler, userHandler, middlewares)

	server.Start()
}
