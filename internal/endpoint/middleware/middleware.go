package middleware

import (
	"context"
	config "hangryAPI/configs"
	"hangryAPI/internal/service/token"
	"net/http"
)

type Func func(handler http.Handler) http.Handler

type contextKey string

func CheckTokenValidity(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cfg := config.NewConfig()
		tokenService := token.NewTokenService(cfg)

		claims, err := tokenService.GetClaims(r, cfg.RefreshSecret)
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), contextKey("claims"), claims)
		rctx := r.WithContext(ctx)

		if claims != nil {
			w.WriteHeader(http.StatusOK)
			next.ServeHTTP(w, rctx)
			return
		}
	})
}
