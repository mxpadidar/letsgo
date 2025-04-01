package services

import (
	"context"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/mxpadidar/letsgo/internal/domain/entities"
	"github.com/mxpadidar/letsgo/internal/domain/errors"
	"github.com/mxpadidar/letsgo/internal/domain/types"
)

type JwtService struct {
	secret []byte
	ttl    int
}

type Claims struct {
	jwt.RegisteredClaims
	UserRole types.Role
}

func NewJwtService(secret []byte, ttl int) *JwtService {
	return &JwtService{secret: secret, ttl: ttl}
}

func (jwtServ *JwtService) Encode(ctx context.Context, user *entities.User) (*types.Token, error) {
	tokenClaims := &Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   user.Username,
			ExpiresAt: jwt.NewNumericDate(jwtServ.getExp()),
			Issuer:    "letsgo",
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		UserRole: user.Role,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims)
	access, err := token.SignedString(jwtServ.secret)
	if err != nil {
		return nil, err
	}

	return types.NewToken(access), nil
}

func (jwtServ *JwtService) Decode(ctx context.Context, tokenString string) (*types.AuthUser, error) {
	authFieldErr := errors.NewAuthFailedError("authentication field!")
	if tokenString == "" {
		return nil, authFieldErr
	}

	customClaims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, customClaims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			errMsg := fmt.Sprintf("unexpected signing method: %s", token.Header["alg"])
			return nil, errors.NewInternalError(errMsg)
		}
		return jwtServ.secret, nil
	})

	if err != nil || !token.Valid {
		return nil, authFieldErr
	}

	// The claims are already of type *TokenClaims
	claims, ok := token.Claims.(*Claims)
	if !ok {
		return nil, authFieldErr
	}

	user := types.NewAuthUser(claims.Subject, claims.UserRole)
	return user, nil
}

func (s *JwtService) getExp() time.Time {
	dur := time.Duration(s.ttl) * time.Second
	return time.Now().Add(dur)
}
