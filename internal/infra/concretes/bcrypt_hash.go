package concretes

import (
	"log"

	"github.com/mxpadidar/letsgo/internal/core/types"
	"golang.org/x/crypto/bcrypt"
)

type BcryptService struct {
	cost int
}

func NewBcryptService(cost int) *BcryptService {
	return &BcryptService{cost: cost}
}

func (s *BcryptService) Hash(raw string) (string, error) {
	if raw == "" {
		return "", types.ErrValidation
	}

	bytesRaw := []byte(raw)

	hash, err := bcrypt.GenerateFromPassword(bytesRaw, s.cost)

	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func (s *BcryptService) Verify(hash, raw string) error {

	if raw == "" || hash == "" {
		return types.ErrValidation
	}

	bytesRaw := []byte(raw)
	bytesHash := []byte(hash)

	log.Printf("Hash length: %d", len(hash))
	log.Printf("Hash value: %s", hash)

	err := bcrypt.CompareHashAndPassword(bytesHash, bytesRaw)
	if err != nil {
		log.Printf("error verifying password: %v", err)
		return err
	}

	return nil
}
