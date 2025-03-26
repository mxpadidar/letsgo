package main

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/mxpadidar/letsgo/internal/api/server"
	"github.com/mxpadidar/letsgo/internal/infra/configs"
	"github.com/mxpadidar/letsgo/internal/infra/dbstore"
	"github.com/mxpadidar/letsgo/internal/infra/helpers"
)

func main() {
	if err := helpers.LoadEnvFile(); err != nil {
		log.Fatal(err)
	}

	configs := configs.InitConfigs()
	db := sqlx.MustConnect("postgres", configs.POSTGRES_DSN)
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	userStore := dbstore.NewUserDBStore(db)

	server := server.NewServer(userStore)

	server.Start(":8000")
}
