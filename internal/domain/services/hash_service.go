package services

import (
	"context"
)

type HashService interface {
	Hash(ctx context.Context, raw string) (hashed []byte, err error)
	Compare(ctx context.Context, hashed []byte, raw string) error
}
