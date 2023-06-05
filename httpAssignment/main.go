package main

import (
	"httpA/endpoint/api"
	"httpA/endpoint/router"
	"net/http"
)

func main() {
	router := router.NewRouter(":3000")
	router.GET("/suppliers", api.GetAllSuppliers, nil)
	router.GET("/suppliers/\\d+", func(w http.ResponseWriter, r *http.Request) {

	}, nil)
	router.Start()
}
