package services

import (
	"context"
	"log"

	"github.com/mxpadidar/letsgo/internal/domain/errors"
	"golang.org/x/crypto/bcrypt"
)

type BcryptService struct {
	cost int
}

func NewBcryptService(cost int) *BcryptService {
	return &BcryptService{cost: cost}
}

func (s *BcryptService) Hash(ctx context.Context, raw string) (hashed string, err error) {
	if raw == "" {
		return "", errors.NewErr(errors.ErrValidation, "raw password cannot be empty", nil)
	}

	bytesRaw := []byte(raw)

	hash, err := bcrypt.GenerateFromPassword(bytesRaw, s.cost)

	if err != nil {
		log.Printf("failed to hash password: %v", err)
		return "", err
	}

	return string(hash), nil
}

func (s *BcryptService) Compare(ctx context.Context, hashed, raw string) error {

	if raw == "" || hashed == "" {
		return errors.NewErr(errors.ErrValidation, "raw or hashed password cannot be empty", nil)
	}

	bytesRaw, bytesHashed := []byte(raw), []byte(hashed)

	err := bcrypt.CompareHashAndPassword(bytesHashed, bytesRaw)
	if err != nil {
		log.Printf("error verifying password: %v", err)
		return err
	}

	return nil
}
