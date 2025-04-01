package services

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/mxpadidar/letsgo/internal/domain/errors"
)

type EnvVarService struct{}

func NewEnvVarService() *EnvVarService {
	return &EnvVarService{}
}

func (e *EnvVarService) LoadEnvironmentFile() error {
	if err := godotenv.Load(); err != nil {
		errMsg := fmt.Sprintf("error loading .env file: %s", err.Error())
		return errors.NewInternalError(errMsg)
	}
	return nil
}

func (e *EnvVarService) GetString(key string) (string, error) {
	if value := os.Getenv(key); value != "" {
		return value, nil
	}
	errMsg := fmt.Sprintf("environment variable %s not found", key)
	return "", errors.NewNotFoundError(errMsg)
}

func (e *EnvVarService) GetInt(key string) (int, error) {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue, nil
		}
	}
	errMsg := fmt.Sprintf("environment variable %s not found", key)
	return 0, errors.NewNotFoundError(errMsg)
}

func (e *EnvVarService) GetBytes(key string) ([]byte, error) {
	if value := os.Getenv(key); value != "" {
		return []byte(value), nil
	}
	errMsg := fmt.Sprintf("environment variable %s not found", key)
	return nil, errors.NewNotFoundError(errMsg)
}
