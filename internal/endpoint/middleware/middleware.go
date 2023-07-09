package middleware

import (
	"context"
	config "hangryAPI/configs"
	handler "hangryAPI/internal/endpoint/http/handlers"
	"hangryAPI/internal/service/token"
	"hangryAPI/internal/util"
	"net/http"
)

type Func func(handler http.Handler) http.Handler

func CheckAccessTokenValidity(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cfg := config.NewConfig(false)
		tokenService := token.NewTokenService(cfg)

		claims, err := tokenService.GetClaims(r, cfg.AccessSecret)
		if err != nil {
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
		cfg := config.NewConfig(false)
		tokenService := token.NewTokenService(cfg)

		claims, err := tokenService.GetClaims(r, cfg.RefreshSecret)
		if err != nil {
			http.Error(w, handler.INVALID_CREDENTIALS, http.StatusUnauthorized)
			return
		}

		if claims != nil {
			next.ServeHTTP(w, r)
			return
		}
	})
}
