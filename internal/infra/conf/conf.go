package conf

import (
	"fmt"
	"os"
	"time"

	"github.com/mxpadidar/letsgo/internal/infra/utils"
)

type Conf struct {
	AppName     string
	AppVersion  string
	Debug       bool
	Port        string
	RootDir     string
	ATDur       time.Duration // Access Token Duration
	RTDur       time.Duration // Refresh Token Duration
	TokenSecret []byte
	// Database configs
	pgDb       string
	pgUser     string
	pgPassword string
	pgHost     string
	pgPort     string
}

func (c *Conf) PgDSN() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		c.pgUser, c.pgPassword, c.pgHost, c.pgPort, c.pgDb,
	)
}

func NewConf() *Conf {
	rootDir, _ := os.Getwd()
	secret := utils.GetEnv("APP_SECRET", "secret")

	return &Conf{
		AppName:     utils.GetEnv("APP_NAME", "MyApp"),
		AppVersion:  utils.GetEnv("APP_VERSION", "1.0.0"),
		Debug:       utils.GetEnv("APP_DEBUG", "false") == "true",
		Port:        utils.GetEnv("APP_PORT", "8000"),
		RootDir:     rootDir,
		ATDur:       time.Hour * 1,
		RTDur:       time.Hour * 24,
		TokenSecret: []byte(secret),
		pgDb:        utils.GetEnv("POSTGRES_DB", "postgres"),
		pgUser:      utils.GetEnv("POSTGRES_USER", "postgres"),
		pgPassword:  utils.GetEnv("POSTGRES_PASSWORD", "postgres"),
		pgHost:      utils.GetEnv("POSTGRES_HOST", "localhost"),
		pgPort:      utils.GetEnv("POSTGRES_PORT", "5432"),
	}
}
