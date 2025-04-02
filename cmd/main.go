package main

import (
	_ "github.com/lib/pq"
	"github.com/mxpadidar/letsgo/internal/bootstrap"
)

func main() {
	srv := bootstrap.Bootstrap()
	srv.Start()
}
