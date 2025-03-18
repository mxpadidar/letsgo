package specs

import "github.com/mxpadidar/letsgo/internal/core/dtos"

type TokenService interface {
	GenerateTokenPair(sub string) (tokenParin *dtos.TokenPair, err error)
	Decode(token, tokenType string) (sub string, err error)
	RefreshTokenPair(refreshToken string) (tokenPair *dtos.TokenPair, err error)
}

type PasswordService interface {
	Hash(raw string) (hash string, err error)
	Verify(hash, raw string) error
}
