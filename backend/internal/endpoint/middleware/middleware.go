package middleware

import (
	"context"
	"fmt"
	config "hangryAPI/configs"
	handler "hangryAPI/internal/endpoint/http/handlers"
	"hangryAPI/internal/service/token"
	"hangryAPI/internal/util"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

type Func func(handler http.Handler) http.Handler

func CheckAccessTokenValidity(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cfg := config.NewConfig("configs/.env")
		tokenService := token.NewTokenService(cfg)

		claims, err := tokenService.GetClaims(r, cfg.AccessSecret)
		if err != nil {
			if fmt.Errorf("%w: %w", jwt.ErrTokenInvalidClaims, jwt.ErrTokenExpired).Error() == err.Error() {
				http.Error(w, handler.ACCESS_TOKEN_EXPIRED, http.StatusUnauthorized)
				return
			}
			http.Error(w, handler.INVALID_CREDENTIALS, http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), util.ContextKey("claims"), claims)
		rctx := r.WithContext(ctx)

		if claims != nil {
			next.ServeHTTP(w, rctx)
			return
		}
	})
}

func CheckRefreshTokenValidity(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cfg := config.NewConfig("configs/.env")
		tokenService := token.NewTokenService(cfg)

		claims, err := tokenService.GetClaims(r, cfg.RefreshSecret)
		if err != nil {
			if fmt.Errorf("%w: %w", jwt.ErrTokenInvalidClaims, jwt.ErrTokenExpired).Error() == err.Error() {
				http.Error(w, handler.REFRESH_TOKEN_EXPIRED, http.StatusUnauthorized)
				return
			}
			http.Error(w, handler.INVALID_CREDENTIALS, http.StatusUnauthorized)
			return
		}

		if claims != nil {
			next.ServeHTTP(w, r)
			return
		}
	})
}
