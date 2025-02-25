package conf

import (
	"fmt"
	"os"
)

type Conf struct {
	AppName          string
	AppVersion       string
	Debug            bool
	Port             string
	RootDir          string
	PostgresDB       string
	PostgresUser     string
	PostgresPassword string
	PostgresHost     string
	PostgresPort     string
}

func NewConf() *Conf {
	loadEnvFile()
	rootDir, _ := os.Getwd()
	return &Conf{
		AppName:          getEnv("APP_NAME", "MyApp"),
		AppVersion:       getEnv("APP_VERSION", "1.0.0"),
		Debug:            getEnv("APP_DEBUG", "false") == "true",
		Port:             getEnv("APP_PORT", "8000"),
		RootDir:          rootDir,
		PostgresDB:       getEnv("POSTGRES_DB", "letsgo"),
		PostgresUser:     getEnv("POSTGRES_USER", "postgres"),
		PostgresPassword: getEnv("POSTGRES_PASSWORD", "password"),
		PostgresHost:     getEnv("POSTGRES_HOST", "localhost"),
		PostgresPort:     getEnv("POSTGRES_PORT", "5432"),
	}
}

func (c *Conf) PgConnStr() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		c.PostgresUser, c.PostgresPassword, c.PostgresHost, c.PostgresPort, c.PostgresDB)
}
