package token

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func (s *TokenService) GenerateToken(tokenSecret string, tokenLifetime uint, userID int) (string, error) {
	claims := &JwtCustomClaims{
		userID,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * time.Duration(tokenLifetime))),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(tokenSecret))
}

func (s *TokenService) GenerateAccessToken(userID int) (string, error) {
	return s.GenerateToken(s.cfg.AccessSecret, s.cfg.AccessLifetime, userID)
}

func (s *TokenService) GenerateRefreshToken(userID int) (string, error) {
	return s.GenerateToken(s.cfg.RefreshSecret, s.cfg.RefreshLifetime, userID)
}

func (s *TokenService) GenerateResetToken(userID int) (string, error) {
	return s.GenerateToken(s.cfg.ResetSecret, s.cfg.ResetLifetime, userID)
}
