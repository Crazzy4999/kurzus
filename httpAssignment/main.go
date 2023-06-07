package main

import (
	"httpA/endpoint/api"
	"httpA/endpoint/http/middleware"
	"httpA/endpoint/router"
)

func main() {
	router := router.NewRouter(":3000")
	router.GET("/suppliers", api.GetAllSuppliers, middleware.LogRequest)
	router.GET("/suppliers/\\d+", api.GetSupplier, middleware.LogRequest)
	router.PUT("/suppliers/\\d+", api.UpdateSupplier, middleware.LogRequest)
	router.POST("/orders", api.AddOrder, middleware.LogRequest)
	router.DELETE("/orders/\\d+", api.DeleteOrder, middleware.LogRequest)
	router.GET("/suppliers/refresh", api.GetAllSuppliersRefresh, middleware.LogRequest)
	router.Start()
}
