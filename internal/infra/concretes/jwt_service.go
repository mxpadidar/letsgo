package concretes

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/mxpadidar/letsgo/internal/core/dtos"
	"github.com/mxpadidar/letsgo/internal/core/types"
)

type JwtService struct {
	secret     []byte
	accessDur  time.Duration
	refreshDur time.Duration
}

func NewJwtService(secret []byte, accessDur, refreshDur time.Duration) *JwtService {
	return &JwtService{secret: secret, accessDur: accessDur, refreshDur: refreshDur}
}

func (s *JwtService) GenerateTokenPair(sub string) (*dtos.TokenPair, error) {
	at, err := s.generate(sub, "access", s.accessDur)
	if err != nil {
		return nil, err
	}

	rt, err := s.generate(sub, "refresh", s.refreshDur)
	if err != nil {
		return nil, err
	}

	return &dtos.TokenPair{AccessToken: at, RefreshToken: rt}, nil
}

func (s *JwtService) Decode(tkn, typ string) (string, error) {
	claims, err := s.parse(tkn)
	if err != nil {
		return "", err
	}

	if claims.Type != typ {
		return "", types.ErrUnauthorized
	}

	if claims.ExpiresAt.Before(time.Now()) {
		return "", types.ErrUnauthorized
	}

	if claims.IssuedAt.Before(time.Now().Add(-time.Hour)) {
		return "", types.ErrUnauthorized
	}

	if claims.IssuedAt.After(time.Now().Add(time.Hour)) {
		return "", types.ErrUnauthorized
	}

	return claims.Subject, nil
}

func (s *JwtService) RefreshTokenPair(rt string) (*dtos.TokenPair, error) {
	sub, err := s.Decode(rt, "refresh")
	if err != nil {
		return nil, err
	}

	return s.GenerateTokenPair(sub)
}

func (s *JwtService) generate(sub, typ string, dur time.Duration) (string, error) {
	now := time.Now()
	exp := now.Add(dur)

	claims := jwt.MapClaims{
		"sub": sub,
		"exp": exp.Unix(),
		"iat": now.Unix(),
		"nbf": now.Unix(),
		"typ": typ,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(s.secret)
}

func (s *JwtService) parse(tknStr string) (*dtos.TokenClaims, error) {
	// keyFunc is a callback function that retrieves the key used to sign the token
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, types.ErrUnauthorized
		}
		return s.secret, nil
	}

	token, err := jwt.Parse(tknStr, keyFunc)
	if err != nil {
		return nil, err
	}

	// Check if the token is valid
	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	// Extract and map the claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, types.ErrUnauthorized
	}

	// Parse token claims
	tc := &dtos.TokenClaims{
		Type:    claims["typ"].(string),
		Subject: claims["sub"].(string),
	}

	if iat, ok := claims["iat"].(float64); ok {
		intTs := int64(iat)
		tc.IssuedAt = time.Unix(intTs, 0)
	} else {
		return nil, types.ErrUnauthorized
	}

	if exp, ok := claims["exp"].(float64); ok {
		intTs := int64(exp)
		tc.ExpiresAt = time.Unix(intTs, 0)
	} else {
		return nil, types.ErrUnauthorized
	}

	if nbf, ok := claims["nbf"].(float64); ok {
		intTs := int64(nbf)
		tc.NotBefore = time.Unix(intTs, 0)
	} else {
		return nil, types.ErrUnauthorized
	}

	return tc, nil
}
