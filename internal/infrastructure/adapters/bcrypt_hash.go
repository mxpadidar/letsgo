package adapters

import (
	"context"

	"github.com/mxpadidar/letsgo/internal/core/errors"
	"github.com/mxpadidar/letsgo/internal/core/services"
	"golang.org/x/crypto/bcrypt"
)

type BcryptHash struct {
	logger services.LogService
}

func NewBcryptHash(logger services.LogService) *BcryptHash {
	return &BcryptHash{logger: logger}
}

func (a *BcryptHash) Hash(ctx context.Context, raw string) (hashed []byte, err error) {
	if raw == "" {
		return nil, errors.NewValidationErr("password is required")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(raw), bcrypt.DefaultCost)

	if err != nil {
		a.logger.Errorf("failed to hash password: %v", err)
		return nil, err
	}

	return hash, nil
}

func (a *BcryptHash) Compare(ctx context.Context, hashed []byte, raw string) error {

	if raw == "" {
		return errors.NewValidationErr("password is required")
	}

	if err := bcrypt.CompareHashAndPassword(hashed, []byte(raw)); err != nil {
		return errors.NewValidationErr("password mismatch")
	}

	return nil
}
