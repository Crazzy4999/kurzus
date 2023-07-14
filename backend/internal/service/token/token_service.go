package token

import (
	"errors"
	config "hangryAPI/configs"
	"hangryAPI/internal/util"
	"net/http"
	"strings"
)

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

func (s *TokenService) GetClaims(r *http.Request, secret string) (*util.JwtCustomClaims, error) {
	authHeader := r.Header.Get("Authorization")
	tokenString := GetTokenFromBearerString(authHeader)

	claims, err := s.ValidateToken(tokenString, secret)
	return claims, err
}

func GetClaimsFromContext(r *http.Request) (*util.JwtCustomClaims, error) {
	claims, ok := r.Context().Value(util.ContextKey("claims")).(*util.JwtCustomClaims)
	if !ok {
		return nil, errors.New("type assertion failed")
	}
	return claims, nil
}
