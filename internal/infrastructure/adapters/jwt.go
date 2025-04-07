package adapters

import (
	"context"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/mxpadidar/letsgo/internal/core/entities"
	"github.com/mxpadidar/letsgo/internal/core/errors"
	"github.com/mxpadidar/letsgo/internal/core/types"
)

type JwtService struct {
	accessSecret  []byte
	refreshSecret []byte
	accessTtl     int
	refreshTtl    int
}

// subject is permit id
type CustomClaims struct {
	jwt.RegisteredClaims
	Role types.Role
}

func NewJwtService(accessSecret, refreshSecret string, accessTtl, refreshTtl int) *JwtService {
	return &JwtService{
		accessSecret:  []byte(accessSecret),
		refreshSecret: []byte(refreshSecret),
		accessTtl:     accessTtl,
		refreshTtl:    refreshTtl,
	}
}

func (s *JwtService) GenerateTokenPair(ctx context.Context, permit *entities.Permit) (*types.TokenPair, error) {
	refreshToken, err := s.generateToken(permit, s.refreshSecret, s.refreshTtl)
	if err != nil || refreshToken == "" {
		return nil, err
	}

	accessToken, err := s.generateToken(permit, s.accessSecret, s.accessTtl)
	if err != nil || accessToken == "" {
		return nil, err
	}

	return types.NewTokenPair(accessToken, refreshToken), nil
}

func (jwtServ *JwtService) DecodeAccessToken(ctx context.Context, tokenString string) (*entities.Permit, error) {
	return jwtServ.decodeToken(tokenString, jwtServ.accessSecret)
}

func (jwtServ *JwtService) DecodeRefreshToken(ctx context.Context, tokenString string) (*entities.Permit, error) {
	return jwtServ.decodeToken(tokenString, jwtServ.refreshSecret)
}

func (jwtServ *JwtService) decodeToken(tokenString string, secret []byte) (*entities.Permit, error) {
	if tokenString == "" {
		return nil, errors.AuthErr
	}

	keyFunc := func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.NewInternalErr("unexpected signing method: %s", token.Header["alg"])
		}
		return secret, nil
	}

	cc := &CustomClaims{}
	token, err := jwt.ParseWithClaims(tokenString, cc, keyFunc)

	if err != nil || !token.Valid {
		return nil, errors.AuthErr
	}

	cc, ok := token.Claims.(*CustomClaims)
	if !ok {
		return nil, errors.AuthErr
	}

	jti, err := uuid.Parse(cc.ID)
	if err != nil {
		return nil, errors.NewInternalErr("invalid UUID: %s", cc.ID)
	}
	userID, err := strconv.Atoi(cc.Subject)
	if err != nil {
		return nil, errors.NewInternalErr("invalid user ID: %s", cc.Subject)
	}
	return entities.NewPermit(jti, userID, cc.Role, cc.IssuedAt.Time), nil
}

func (jwtServ *JwtService) generateToken(permit *entities.Permit, secret []byte, ttl int) (string, error) {
	now := time.Now()
	exp := now.Add(time.Duration(ttl) * time.Second)
	cc := CustomClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        permit.ID.String(),
			Subject:   strconv.Itoa(permit.UserID),
			ExpiresAt: jwt.NewNumericDate(exp),
			IssuedAt:  jwt.NewNumericDate(now),
		},
		Role: permit.Role,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, cc)
	return token.SignedString(secret)
}
