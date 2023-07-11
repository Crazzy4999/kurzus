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
	cfg := config.NewConfig(".env")
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
	router.GET("/refresh", authHandler.Refresh, middleware.CheckRefreshTokenValidity)

	router.GET("/profile", userHandler.GetProfile, middleware.CheckAccessTokenValidity)
	router.PUT("/profile", userHandler.UpdateProfile, middleware.CheckAccessTokenValidity)
	router.DELETE("/profile", userHandler.DeleteProfile, middleware.CheckAccessTokenValidity)

	router.POST("/address", addressHandler.AddAddress, middleware.CheckAccessTokenValidity)
	router.GET("/addresses", addressHandler.GetAddressesByUserID, middleware.CheckAccessTokenValidity)
	router.PUT("/address", addressHandler.UpdateAddress, middleware.CheckAccessTokenValidity)
	router.DELETE("/address", addressHandler.RemoveAddress, middleware.CheckAccessTokenValidity)

	router.POST("/driver", driverHandler.AddDriver, middleware.CheckAccessTokenValidity)
	router.GET("/drivers", driverHandler.GetDrivers, middleware.CheckAccessTokenValidity)
	router.PUT("/driver", driverHandler.UpdateDriver, middleware.CheckAccessTokenValidity)
	router.DELETE("/driver", driverHandler.RemoveDriver, middleware.CheckAccessTokenValidity)

	router.POST("/supplier", supplierHandler.AddSupplier, middleware.CheckAccessTokenValidity)
	router.GET("/supplier/\\d+", supplierHandler.GetSupplierByID, nil)
	router.GET("/suppliers", supplierHandler.GetSuppliers, nil)
	router.PUT("/supplier", supplierHandler.UpdateSupplier, middleware.CheckAccessTokenValidity)
	router.DELETE("/supplier", supplierHandler.RemoveSupplier, middleware.CheckAccessTokenValidity)

	router.POST("/category", categoryHandler.AddCategory, middleware.CheckAccessTokenValidity)
	router.GET("/categories", categoryHandler.GetCategories, nil)
	router.DELETE("/category", categoryHandler.RemoveCategory, middleware.CheckAccessTokenValidity)

	router.POST("/item", itemHandler.AddItem, middleware.CheckAccessTokenValidity)
	router.GET("/items", itemHandler.GetItems, nil)
	router.DELETE("/item", itemHandler.RemoveItem, middleware.CheckAccessTokenValidity)

	router.POST("/menu", menuHandler.AddMenu, middleware.CheckAccessTokenValidity)
	router.GET("/menus", menuHandler.GetMenus, nil)
	router.DELETE("/menu", menuHandler.RemoveMenu, middleware.CheckAccessTokenValidity)

	router.GET("/itemsmenus/\\d+", nil, nil)

	router.POST("/ordermenu", orderMenuHandler.AddOrderMenu, middleware.CheckAccessTokenValidity)
	router.GET("/ordermenus/\\d+", orderMenuHandler.GetOrderMenus, middleware.CheckAccessTokenValidity)
	router.PUT("/ordermenu", orderMenuHandler.UpdateOrderMenu, middleware.CheckAccessTokenValidity)
	router.DELETE("/ordermenu", orderMenuHandler.RemoveOrderMenu, middleware.CheckAccessTokenValidity)

	router.POST("/order", orderHandler.MakeOrder, middleware.CheckAccessTokenValidity)
	router.GET("/orders", orderHandler.GetOrders, middleware.CheckAccessTokenValidity)
	router.PUT("/order", orderHandler.UpdateOrder, middleware.CheckAccessTokenValidity)

	http.ListenAndServe(cfg.Port, router)
}
