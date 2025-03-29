package services

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/mxpadidar/letsgo/internal/domain/dtos"
	"github.com/mxpadidar/letsgo/internal/domain/entities"
	"github.com/mxpadidar/letsgo/internal/domain/errors"
)

type JwtService struct {
	secret []byte
	ttl    int
}

func NewJwtService(secret []byte, ttl int) *JwtService {
	return &JwtService{secret: secret, ttl: ttl}
}

func (jwtServ *JwtService) Encode(ctx context.Context, user *entities.User) (*dtos.TokenDTO, error) {
	claims := jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Duration(jwtServ.ttl) * time.Second).Unix(),
		"iss": "letsgo",
		"iat": time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	access, err := token.SignedString(jwtServ.secret)
	if err != nil {
		return nil, err
	}

	return dtos.NewTokenDTO(access), nil
}

func (jwtServ *JwtService) Decode(ctx context.Context, tokenString string) (payload *dtos.TokenPayload, err error) {

	if tokenString == "" {
		return nil, errors.NewErr(errors.ErrAuthFailed, "token is empty", nil)
	}

	tokenClaims, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.NewErr(errors.ErrAuthFailed, "unexpected signing method", nil)
		}
		return jwtServ.secret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := tokenClaims.Claims.(jwt.MapClaims); ok && tokenClaims.Valid {
		// Safer approach with error handling
		if sub, ok := claims["sub"].(float64); ok {
			log.Printf("sub: %v", sub)
			userID := int(sub)
			payload = dtos.NewTokenPayload(userID)
			return payload, nil
		}
		return nil, fmt.Errorf("invalid claim type for 'sub'")
	}

	return nil, errors.NewErr(errors.ErrAuthFailed, "invalid token", nil)
}
