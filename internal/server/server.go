package server

import (
	"context"
	"fmt"
	config "hangryAPI/configs"
	handler "hangryAPI/internal/endpoint/http/handlers"
	"hangryAPI/internal/endpoint/middleware"
	dbrepo "hangryAPI/internal/repositories/db"
	"hangryAPI/internal/router"
	"hangryAPI/internal/workers"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func Start() {
	cfg := config.NewConfig("configs/.env")
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
	itemsMenusHandler := handler.NewItemsMenusHandler(imr)
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
	router.GET("/address/\\d+", addressHandler.GetAddressByID, middleware.CheckAccessTokenValidity)
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
	router.GET("/menus/\\d+", menuHandler.GetMenusBySupplierID, nil)
	router.GET("/menus", menuHandler.GetMenus, nil)
	router.DELETE("/menu", menuHandler.RemoveMenu, middleware.CheckAccessTokenValidity)

	router.GET("/itemsmenus/\\d+", itemsMenusHandler.GetItemsByMenuID, nil)

	router.POST("/ordermenu", orderMenuHandler.AddOrderMenu, middleware.CheckAccessTokenValidity)
	router.GET("/ordermenu/\\d+", orderMenuHandler.GetMenusByOrderID, middleware.CheckAccessTokenValidity)
	router.GET("/ordermenus/\\d+", orderMenuHandler.GetOrderMenus, middleware.CheckAccessTokenValidity)
	router.PUT("/ordermenu", orderMenuHandler.UpdateOrderMenu, middleware.CheckAccessTokenValidity)
	router.DELETE("/ordermenu", orderMenuHandler.RemoveOrderMenu, middleware.CheckAccessTokenValidity)

	router.POST("/order", orderHandler.MakeOrder, middleware.CheckAccessTokenValidity)
	router.GET("/orders", orderHandler.GetOrders, middleware.CheckAccessTokenValidity)
	router.PUT("/order", orderHandler.UpdateOrder, middleware.CheckAccessTokenValidity)

	orderWorker := workers.NewOrderDriverWorkerPool(or, dr)
	go orderWorker.Start(5)

	deliveryWorker := workers.NewOrderDeliveryWorkerPool(or, dr)
	go deliveryWorker.Start(5)

	srv := &http.Server{Handler: router}
	ln, err := net.Listen("tcp", "localhost"+cfg.Port)
	if err != nil {
		panic("failed init listener")
	}

	go func() {
		fmt.Printf("Server started on port: %s\n", cfg.Port)

		log.Fatal(srv.ServeTLS(ln,
			"localhost.pem",
			"localhost-key.pem",
		))
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	fmt.Println("Gracefully stopping the server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}
