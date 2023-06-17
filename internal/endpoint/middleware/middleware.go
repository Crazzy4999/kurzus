package middleware

import (
	config "hangryAPI/configs"
	"hangryAPI/internal/service/token"
	"net/http"
)

type Func func(handler http.Handler) http.Handler

func CheckAccessTokenValidity(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cfg := config.NewConfig()
		tokenService := token.NewTokenService(cfg)

		authHeader := r.Header.Get("Authorization")
		tokenString := token.GetTokenFromBearerString(authHeader)

		claims, err := tokenService.ValidateAccessToken(tokenString)
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusUnauthorized)
			return
		}

		if claims != nil {
			w.WriteHeader(http.StatusOK)
			next.ServeHTTP(w, r)
			return
		}
	})
}
