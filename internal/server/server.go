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
	defer db.Close()

	/*dr := dbrepo.NewDriverRepository(db)
	or := dbrepo.NewOrderRepository(db)
	sr := dbrepo.NewSupplierRepository(db)*/
	ur := dbrepo.NewUserRepository(db)
	ar := dbrepo.NewAddressRepository(db)

	authHandler := handler.NewAuthHandler(ur, cfg)
	userHandler := handler.NewUserHandler(ur, cfg)
	addressHandler := handler.NewAddressHandler(ar, cfg)

	router.GET("/", nil, nil)

	router.POST("/login", authHandler.Login, nil)
	router.POST("/signup", authHandler.SignUp, nil)
	router.POST("/reset", authHandler.GetPasswordResetKey, nil)
	router.PUT("/reset\\?reset_key=.+", authHandler.ResetPassword, nil)
	router.GET("/refresh", authHandler.Refresh, middleware.CheckTokenValidity)

	router.GET("/profile", userHandler.GetProfile, middleware.CheckTokenValidity)
	router.PUT("/profile", userHandler.UpdateProfile, middleware.CheckTokenValidity)

	router.POST("/address", addressHandler.AddAddress, middleware.CheckTokenValidity)
	//router.GET("/address", addressHandler.GetAddresses, middleware.CheckTokenValidity)
	router.PUT("/address", addressHandler.UpdateAddress, middleware.CheckTokenValidity)
	//router.DELETE("/address", addressHandler.RemoveAddress, middleware.CheckTokenValidity)

	router.POST("/suppliers", nil, middleware.CheckTokenValidity) //Create new supplier
	router.GET("/suppliers", nil, middleware.CheckTokenValidity)
	router.GET("/suppliers/\\d+", nil, middleware.CheckTokenValidity)
	router.GET("/categories", nil, middleware.CheckTokenValidity)
	router.GET("/cart", nil, middleware.CheckTokenValidity)
	router.POST("/order", nil, middleware.CheckTokenValidity)
	router.GET("/history", nil, middleware.CheckTokenValidity)

	http.ListenAndServe(cfg.Port, router)
}
