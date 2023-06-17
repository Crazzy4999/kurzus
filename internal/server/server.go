package server

import (
	config "hangryAPI/configs"
	"hangryAPI/internal/endpoint/middleware"
	"hangryAPI/internal/repositories/db"
	"hangryAPI/internal/router"
	"net/http"
)

func Start() {
	cfg := config.NewConfig()
	router := router.NewRouter()
	_ = db.GetDB()

	router.GET("/", nil, nil)
	router.POST("/login", nil, nil)
	router.POST("/signup", nil, nil)
	router.GET("/refresh", nil, middleware.CheckAccessTokenValidity)
	router.GET("/suppliers", nil, middleware.CheckAccessTokenValidity)
	router.GET("/suppliers/\\d+", nil, middleware.CheckAccessTokenValidity)
	router.GET("/categories", nil, middleware.CheckAccessTokenValidity)
	router.GET("/cart", nil, middleware.CheckAccessTokenValidity)
	router.POST("/order", nil, middleware.CheckAccessTokenValidity)
	router.GET("/history", nil, middleware.CheckAccessTokenValidity)

	http.ListenAndServe(cfg.Port, router)
}
