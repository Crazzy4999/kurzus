package server

import (
	config "hangryAPI/configs"
	handler "hangryAPI/internal/endpoint/http/handlers"
	"hangryAPI/internal/endpoint/middleware"
	dbrepo "hangryAPI/internal/repositories/db"
	"hangryAPI/internal/router"
	"net/http"
)

func Start() {
	cfg := config.NewConfig()
	router := router.NewRouter()
	db := dbrepo.GetDB()

	/*dr := dbrepo.NewDriverRepository(db)
	or := dbrepo.NewOrderRepository(db)
	sr := dbrepo.NewSupplierRepository(db)*/
	ur := dbrepo.NewUserRepository(db)

	authHandler := handler.NewAuthHandler(ur, cfg)

	router.GET("/", nil, nil)
	router.POST("/login", authHandler.Login, nil)
	router.POST("/signup", authHandler.SignUp, nil)
	router.GET("/refresh", authHandler.Refresh, middleware.CheckAccessTokenValidity)
	router.GET("/suppliers", nil, middleware.CheckAccessTokenValidity)
	router.GET("/suppliers/\\d+", nil, middleware.CheckAccessTokenValidity)
	router.GET("/categories", nil, middleware.CheckAccessTokenValidity)
	router.GET("/cart", nil, middleware.CheckAccessTokenValidity)
	router.POST("/order", nil, middleware.CheckAccessTokenValidity)
	router.GET("/history", nil, middleware.CheckAccessTokenValidity)

	http.ListenAndServe(cfg.Port, router)
}
