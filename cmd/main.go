package main

import (
	"fmt"

	"github.com/mxpadidar/letsgo/internal/api/middlewares"
	"github.com/mxpadidar/letsgo/internal/api/server"
	"github.com/mxpadidar/letsgo/internal/infra/concretes"
	"github.com/mxpadidar/letsgo/internal/infra/conf"
	"github.com/mxpadidar/letsgo/internal/infra/db"
	"github.com/mxpadidar/letsgo/internal/infra/logger"
	"github.com/mxpadidar/letsgo/internal/infra/pgstores"
	"github.com/mxpadidar/letsgo/internal/infra/utils"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	utils.LoadEnvFile()
	configs := conf.NewConf()
	log := logger.New("main")

	pg, err := db.NewPgDb(configs.PgDSN())
	if err != nil {
		log.Fatal("failed to connect to database", err)
	}

	userPgStore := pgstores.NewUserPgStore(pg.Db)

	bcryptService := concretes.NewBcryptService(bcrypt.DefaultCost)

	jwtService := concretes.NewJwtService(configs.TokenSecret, configs.ATDur, configs.RTDur)

	addr := fmt.Sprintf(":%s", configs.Port)
	auth_mw := middlewares.NewAuthMiddleware(jwtService, userPgStore)
	mw := middlewares.MiddlewareChain(middlewares.LogMiddleware, auth_mw)
	server := server.NewServer(addr, jwtService, bcryptService, userPgStore, mw)
	server.Setup()

	log.Info("%s is starting", configs.AppName)

	if err := server.Start(); err != nil {
		log.Fatal("failed to start server", err)
	}
}
