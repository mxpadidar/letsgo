package servicesimpl

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type BcryptHashService struct{}

func NewBcryptHashService() *BcryptHashService {
    return &BcryptHashService{}
}

func (h *BcryptHashService) HashPassword(password string) (string, error) {
    if password == "" {
        return "", fmt.Errorf("password cannot be empty")
    }

    hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return "", fmt.Errorf("failed to hash password: %w", err)
    }

    return string(hash), nil
}

func (h *BcryptHashService) ComparePassword(hashPassword, password string) bool {
    if hashPassword == "" || password == "" {
        return false
    }

    err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
    return err == nil
}
