package token

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func (s *TokenService) GenerateToken(tokenSecret string, tokenLifetime uint, userId int) (string, error) {
	claims := &JwtCustomClaims{
		userId,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * time.Duration(tokenLifetime))),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(tokenSecret))
}

func (s *TokenService) GenerateAccessToken(userId int) (string, error) {
	return s.GenerateToken(s.cfg.AccessSecret, s.cfg.AccessLifetime, userId)
}

func (s *TokenService) GenerateRefreshToken(userId int) (string, error) {
	return s.GenerateToken(s.cfg.RefreshSecret, s.cfg.RefreshLifetime, userId)
}
