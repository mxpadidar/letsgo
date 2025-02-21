package main

import (
	"github.com/mxpadidar/letsgo/pkg/conf"
	"github.com/mxpadidar/letsgo/pkg/logger"
)

func main() {
	configs := conf.NewConf()
	log := logger.New("main")

	log.Info("%s is starting", configs.AppName)
}
