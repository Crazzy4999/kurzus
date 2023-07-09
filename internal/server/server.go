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
	cfg := config.NewConfig(false)
	router := router.NewRouter()
	db := dbrepo.GetDB(cfg)
	defer db.Close()

	ur := dbrepo.NewUserRepository(db)
	ar := dbrepo.NewAddressRepository(db)
	dr := dbrepo.NewDriverRepository(db)
	sr := dbrepo.NewSupplierRepository(db)
	str := dbrepo.NewSupplierTypesRepository(db)
	or := dbrepo.NewOrderRepository(db)
	osr := dbrepo.NewOrderStatusRepository(db)
	cr := dbrepo.NewCategoriesRepository(db)
	ir := dbrepo.NewItemRepository(db)
	imr := dbrepo.NewItemsMenusRepository(db)
	mr := dbrepo.NewMenuRepository(db)
	omr := dbrepo.NewOrderMenusRepository(db)

	authHandler := handler.NewAuthHandler(ur, cfg)
	userHandler := handler.NewUserHandler(ur, ar, cfg)
	addressHandler := handler.NewAddressHandler(ar)
	driverHandler := handler.NewDriverHandler(dr)
	supplierHandler := handler.NewSupplierHandler(sr, str)
	orderHandler := handler.NewOrderHandler(or, osr)
	categoryHandler := handler.NewCategoryHandler(cr)
	itemHandler := handler.NewItemHandler(ir, imr)
	menuHandler := handler.NewMenuHandler(mr, imr)
	orderMenuHandler := handler.NewOrderMenuHandler(omr)

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
	router.GET("/addresses", addressHandler.GetAddressesByUserID, middleware.CheckTokenValidity)
	router.PUT("/address", addressHandler.UpdateAddress, middleware.CheckTokenValidity)
	router.DELETE("/address", addressHandler.RemoveAddress, middleware.CheckTokenValidity)

	router.POST("/driver", driverHandler.AddDriver, middleware.CheckTokenValidity)
	router.GET("/drivers", driverHandler.GetDrivers, middleware.CheckTokenValidity)
	router.PUT("/driver", driverHandler.UpdateDriver, middleware.CheckTokenValidity)
	router.DELETE("/driver", driverHandler.RemoveDriver, middleware.CheckTokenValidity)

	router.POST("/supplier", supplierHandler.AddSupplier, middleware.CheckTokenValidity)
	router.GET("/suppliers", supplierHandler.GetSuppliers, middleware.CheckTokenValidity)
	router.PUT("/supplier", supplierHandler.UpdateSupplier, middleware.CheckTokenValidity)
	router.DELETE("/supplier", supplierHandler.RemoveSupplier, middleware.CheckTokenValidity)

	router.POST("/category", categoryHandler.AddCategory, middleware.CheckTokenValidity)
	router.GET("/categories", categoryHandler.GetCategories, middleware.CheckTokenValidity)
	router.DELETE("/category", categoryHandler.RemoveCategory, middleware.CheckTokenValidity)

	router.POST("/item", itemHandler.AddItem, middleware.CheckTokenValidity)
	router.GET("/items", itemHandler.GetItems, middleware.CheckTokenValidity)
	router.DELETE("/item", itemHandler.RemoveItem, middleware.CheckTokenValidity)

	router.POST("/menu", menuHandler.AddMenu, middleware.CheckTokenValidity)
	router.GET("/menus", menuHandler.GetMenus, middleware.CheckTokenValidity)
	router.DELETE("/menu", menuHandler.RemoveMenu, middleware.CheckTokenValidity)

	router.POST("/ordermenu", orderMenuHandler.AddOrderMenu, middleware.CheckTokenValidity)
	router.GET("/ordermenus", orderMenuHandler.GetOrderMenus, middleware.CheckTokenValidity)
	router.PUT("/ordermenu", orderMenuHandler.UpdateOrderMenu, middleware.CheckTokenValidity)
	router.DELETE("/ordermenu", orderMenuHandler.RemoveOrderMenu, middleware.CheckTokenValidity)

	router.POST("/order", orderHandler.MakeOrder, middleware.CheckTokenValidity)
	router.GET("/orders", orderHandler.GetOrders, middleware.CheckTokenValidity)
	router.PUT("/order", orderHandler.UpdateOrder, middleware.CheckTokenValidity)

	http.ListenAndServe(cfg.Port, router)
}
