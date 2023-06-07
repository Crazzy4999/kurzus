package api

import (
	"httpA/db/model"
	"httpA/db/repo/file"
	"httpA/endpoint/http/domain"
	"httpA/endpoint/http/request"
	"httpA/endpoint/http/response"
	"net/http"
	"regexp"
	"strconv"
)

func AddOrder(w http.ResponseWriter, r *http.Request) {
	cro := request.CreateOrderRequest{}
	or := &file.OrdersRepository{}

	err := request.ParseJsonRequest(r, &cro)
	if err != nil {
		response.SendBadRequestError(w, err)
		return
	}

	count := or.OrdersCount()
	order := model.Order{OrderId: count, SupplierId: cro.SupplierId, UserId: cro.UserId, Address: struct {
		Street   string "json:\"street\""
		Building string "json:\"building\""
		Apt      int    "json:\"apt\""
	}{Street: cro.Address.Street, Building: cro.Address.Building, Apt: cro.Address.Apt}, Status: cro.Status}

	domain.AddOrder(&order)
	response.SendOK(w, order)
}

func DeleteOrder(w http.ResponseWriter, r *http.Request) {
	regex := regexp.MustCompile("\\d+")
	id := regex.FindString(r.URL.Path)
	intId, err := strconv.Atoi(id)
	if err != nil {
		response.SendBadRequestError(w, err)
		return
	}

	domain.DeleteOrder(intId)
}
