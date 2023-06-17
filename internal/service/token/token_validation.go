package token

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
)

func (s *TokenService) ValidateToken(tokenString, secret string) (*JwtCustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JwtCustomClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*JwtCustomClaims)
	if !ok || !token.Valid {
		return nil, errors.New("failed to parse token claims")
	}

	return claims, nil
}

func (s *TokenService) ValidateAccessToken(tokenString string) (*JwtCustomClaims, error) {
	return s.ValidateToken(tokenString, s.cfg.AccessSecret)
}

func (s *TokenService) ValidateRefreshToken(tokenString string) (*JwtCustomClaims, error) {
	return s.ValidateToken(tokenString, s.cfg.RefreshSecret)
}
