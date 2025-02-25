package main

import (
	"fmt"

	"github.com/mxpadidar/letsgo/internal/api/server"
	"github.com/mxpadidar/letsgo/internal/infra/db"
	"github.com/mxpadidar/letsgo/internal/infra/pgstores"
	"github.com/mxpadidar/letsgo/internal/infra/servicesimpl"
	"github.com/mxpadidar/letsgo/pkg/conf"
	"github.com/mxpadidar/letsgo/pkg/logger"
)

func main() {
	configs := conf.NewConf()
	log := logger.New("main")

	pg, err := db.NewPgDb(configs.PgConnStr())
	if err != nil {
		log.Fatal("failed to connect to database", err)
	}

	userStore := pgstores.NewUserPgStore(pg.Db)

	hash := servicesimpl.NewBcryptHashService()
	addr := fmt.Sprintf(":%s", configs.Port)
	server := server.NewServer(addr, userStore, hash)
	server.SetupHandlers()

	log.Info("%s is starting", configs.AppName)

	if err := server.Start(); err != nil {
		log.Fatal("failed to start server", err)
	}
}
