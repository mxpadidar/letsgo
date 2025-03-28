package helpers

import (
	"log"
	"os"
	"strconv"
)

func MustEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("environment variable %s is not set", key)
	}
	return value
}


func MustIntEnv(key string) int {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("environment variable %s is not set", key)
	}
	intValue, err := strconv.Atoi(value)
	if err != nil {
		log.Fatalf("environment variable %s is not a valid integer", key)
	}
	return intValue
}
