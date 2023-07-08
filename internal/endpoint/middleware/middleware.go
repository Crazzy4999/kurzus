package middleware

import (
	"context"
	config "hangryAPI/configs"
	"hangryAPI/internal/service/token"
	"hangryAPI/internal/util"
	"net/http"
)

type Func func(handler http.Handler) http.Handler

func CheckTokenValidity(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cfg := config.NewConfig(false)
		tokenService := token.NewTokenService(cfg)

		claims, err := tokenService.GetClaims(r, cfg.RefreshSecret)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
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
