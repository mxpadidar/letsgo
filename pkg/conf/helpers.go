package conf

import (
	"os"

	"github.com/joho/godotenv"
)

func loadEnvFile() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}

func getEnv(key, fallback string) string {
	env := os.Getenv(key)
	if env == "" {
		return fallback
	}
	return env
}
