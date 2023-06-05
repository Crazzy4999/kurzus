package api

import (
	"httpA/endpoint/http/domain"
	"httpA/endpoint/http/response"
	"net/http"
)

func GetAllSuppliers(w http.ResponseWriter, r *http.Request) {
	response.SendOK(w, domain.GetAllSuppliers)
}
