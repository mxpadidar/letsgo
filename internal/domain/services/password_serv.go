package services

import "context"

type PasswordService interface {
	// generate hashed password from raw password
	Hash(ctx context.Context, raw string) (hashed string, err error)

	// compare hashed password with raw password, return error if not match
	Compare(ctx context.Context, hashed, raw string) error
}
