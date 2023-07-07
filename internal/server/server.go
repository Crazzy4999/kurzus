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

	//dr := dbrepo.NewDriverRepository(db)
	ur := dbrepo.NewUserRepository(db)
	ar := dbrepo.NewAddressRepository(db)
	sr := dbrepo.NewSupplierRepository(db)
	str := dbrepo.NewSupplierTypesRepository(db)
	or := dbrepo.NewOrderRepository(db)
	omr := dbrepo.NewOrderMenusRepository(db)
	osr := dbrepo.NewOrderStatusRepository(db)
	cr := dbrepo.NewCategoriesRepository(db)
	ir := dbrepo.NewItemRepository(db)

	authHandler := handler.NewAuthHandler(ur, cfg)
	userHandler := handler.NewUserHandler(ur, ar, cfg)
	addressHandler := handler.NewAddressHandler(ar)
	supplierHandler := handler.NewSupplierHandler(sr, str)
	orderHandler := handler.NewOrderHandler(or, omr, osr)
	categoryHandler := handler.NewCategoryHandler(cr)
	itemHandler := handler.NewItemHandler(ir)

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

	router.POST("/categories", categoryHandler.AddCategory, middleware.CheckTokenValidity)
	router.GET("/categories", categoryHandler.GetCategories, middleware.CheckTokenValidity)
	router.DELETE("/categories", categoryHandler.RemoveCategory, middleware.CheckTokenValidity)

	router.POST("/item", itemHandler.AddItem, middleware.CheckTokenValidity)
	router.GET("/item", itemHandler.GetItems, middleware.CheckTokenValidity)
	router.DELETE("/item", itemHandler.RemoveItem, middleware.CheckTokenValidity)

	router.POST("/menu", nil, middleware.CheckTokenValidity)
	router.GET("/menu", nil, middleware.CheckTokenValidity)
	router.PUT("/menu", nil, middleware.CheckTokenValidity)
	router.DELETE("/menu", nil, middleware.CheckTokenValidity)

	router.POST("/ordermenu", nil, middleware.CheckTokenValidity)
	router.GET("/ordermenu", nil, middleware.CheckTokenValidity)
	router.PUT("/ordermenu", nil, middleware.CheckTokenValidity)
	router.DELETE("/ordermenu", nil, middleware.CheckTokenValidity)

	router.POST("/order", orderHandler.MakeOrder, middleware.CheckTokenValidity)
	router.GET("/orders", orderHandler.GetOrders, middleware.CheckTokenValidity)
	router.PUT("/order", orderHandler.UpdateOrder, middleware.CheckTokenValidity)

	http.ListenAndServe(cfg.Port, router)
}
