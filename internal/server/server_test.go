package server

import (
	"hangryAPI/internal/router"
	"net/http/httptest"
)

func StartTest() *httptest.Server {
	router := router.NewRouter()

	return httptest.NewServer(router)
}
