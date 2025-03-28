package main

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/mxpadidar/letsgo/internal/api/server"
	"github.com/mxpadidar/letsgo/internal/infra/configs"
	"github.com/mxpadidar/letsgo/internal/infra/dbstore"
	"github.com/mxpadidar/letsgo/internal/infra/helpers"
	"github.com/mxpadidar/letsgo/internal/infra/services"
)

func main() {
	if err := helpers.LoadEnvFile(); err != nil {
		log.Fatal(err)
	}

	configs := configs.InitConfigs()
	db := sqlx.MustConnect("postgres", configs.PostgresDSN)
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	userStore := dbstore.NewUserDBStore(db)
	jwtService := services.NewJwtService([]byte(configs.JWTSecret), configs.AccessTokenTTL)
	bcryptService := services.NewBcryptService(configs.BcryptCost)

	server := server.NewServer(userStore, jwtService, bcryptService)

	server.Start(":8000")
}
