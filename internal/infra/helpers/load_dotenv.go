package helpers

import (
	"github.com/joho/godotenv"
	"github.com/mxpadidar/letsgo/internal/domain/errors"
)

func LoadEnvFile() error {
	err := godotenv.Load()
	if err != nil {
		return errors.NewErr(errors.ErrValidation, "Failed to load .env file", err)
	}
	return nil
}
