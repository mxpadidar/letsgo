package utils

import (
	"os"

	"github.com/joho/godotenv"
)

func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func LoadEnvFile() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}
