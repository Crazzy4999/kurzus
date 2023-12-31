package token

import (
	"errors"
	"hangryAPI/internal/util"

	"github.com/golang-jwt/jwt/v5"
)

func (s *TokenService) ValidateToken(tokenString, secret string) (*util.JwtCustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &util.JwtCustomClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*util.JwtCustomClaims)
	if !ok || !token.Valid {
		return nil, errors.New("failed to parse token claims")
	}

	return claims, nil
}

func (s *TokenService) ValidateAccessToken(tokenString string) (*util.JwtCustomClaims, error) {
	return s.ValidateToken(tokenString, s.cfg.AccessSecret)
}

func (s *TokenService) ValidateRefreshToken(tokenString string) (*util.JwtCustomClaims, error) {
	return s.ValidateToken(tokenString, s.cfg.RefreshSecret)
}

func (s *TokenService) ValidateResetToken(tokenString string) (*util.JwtCustomClaims, error) {
	return s.ValidateToken(tokenString, s.cfg.ResetSecret)
}
