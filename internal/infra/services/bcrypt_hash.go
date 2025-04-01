package services

import (
	"context"
	"log"

	"github.com/mxpadidar/letsgo/internal/domain/errors"
	"golang.org/x/crypto/bcrypt"
)

type BcryptHash struct{}

func NewBcryptHash() *BcryptHash {
	return &BcryptHash{}
}

func (a *BcryptHash) Hash(ctx context.Context, raw string) (hashed []byte, err error) {
	if raw == "" {
		return nil, errors.NewValidationError("password is required")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(raw), bcrypt.DefaultCost)

	if err != nil {
		log.Printf("failed to hash password: %v", err)
		return nil, err
	}

	return hash, nil
}

func (a *BcryptHash) Compare(ctx context.Context, hashed []byte, raw string) error {

	if raw == "" {
		return errors.NewValidationError("password is required")
	}

	if err := bcrypt.CompareHashAndPassword(hashed, []byte(raw)); err != nil {
		return errors.NewValidationError("password mismatch")
	}

	return nil
}
