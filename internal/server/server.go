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
	or := dbrepo.NewOrderRepository(db)*/
	ur := dbrepo.NewUserRepository(db)
	ar := dbrepo.NewAddressRepository(db)
	sr := dbrepo.NewSupplierRepository(db)
	str := dbrepo.NewSupplierTypesRepository(db)

	authHandler := handler.NewAuthHandler(ur, cfg)
	userHandler := handler.NewUserHandler(ur, ar, cfg)
	addressHandler := handler.NewAddressHandler(ar)
	supplierHandler := handler.NewSupplierHandler(sr, str)

	router.GET("/", nil, nil)

	router.POST("/login", authHandler.Login, nil)
	router.POST("/signup", authHandler.SignUp, nil)
	router.POST("/reset", authHandler.GetPasswordResetKey, nil)
	router.PUT("/reset\\?reset_key=.+", authHandler.ResetPassword, nil)
	router.GET("/refresh", authHandler.Refresh, middleware.CheckTokenValidity)

	router.GET("/profile", userHandler.GetProfile, middleware.CheckTokenValidity)
	router.PUT("/profile", userHandler.UpdateProfile, middleware.CheckTokenValidity)
	router.DELETE("/profile", userHandler.DeleteProfile, middleware.CheckTokenValidity)

	router.POST("/address", addressHandler.AddAddress, middleware.CheckTokenValidity)
	router.GET("/address", addressHandler.GetAddressesByUserID, middleware.CheckTokenValidity)
	router.PUT("/address", addressHandler.UpdateAddress, middleware.CheckTokenValidity)
	router.DELETE("/address", addressHandler.RemoveAddress, middleware.CheckTokenValidity)

	router.POST("/suppliers", supplierHandler.AddSupplier, middleware.CheckTokenValidity)
	router.GET("/suppliers", supplierHandler.GetSuppliers, middleware.CheckTokenValidity)
	router.PUT("/suppliers", supplierHandler.UpdateSupplier, middleware.CheckTokenValidity)
	router.DELETE("/suppliers", supplierHandler.RemoveSupplier, middleware.CheckTokenValidity)

	router.GET("/categories", nil, middleware.CheckTokenValidity)
	router.GET("/cart", nil, middleware.CheckTokenValidity)
	router.POST("/order", nil, middleware.CheckTokenValidity)
	router.GET("/history", nil, middleware.CheckTokenValidity)

	http.ListenAndServe(cfg.Port, router)
}
