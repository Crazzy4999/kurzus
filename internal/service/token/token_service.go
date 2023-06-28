package token

import (
	config "hangryAPI/configs"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

type JwtCustomClaims struct {
	ID int `json:"id"`
	jwt.RegisteredClaims
}

type TokenService struct {
	cfg *config.Config
}

func NewTokenService(cfg *config.Config) *TokenService {
	return &TokenService{
		cfg: cfg,
	}
}

func GetTokenFromBearerString(bearerString string) string {
	if bearerString == "" {
		return ""
	}

	parts := strings.Split(bearerString, "Bearer")
	if len(parts) != 2 {
		return ""
	}

	token := strings.TrimSpace(parts[1])
	if len(token) < 1 {
		return ""
	}

	return token
}

func (s *TokenService) GetClaims(r *http.Request, secret string) (*JwtCustomClaims, error) {
	authHeader := r.Header.Get("Authorization")
	tokenString := GetTokenFromBearerString(authHeader)

	claims, err := s.ValidateToken(tokenString, secret)
	return claims, err
}
