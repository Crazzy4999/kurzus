package handler

import (
	"hangryAPI/internal/repositories/db"
	"net/http"
)

type OrderHandler struct {
	orderRepo       *db.OrderRepository
	orderMenusRepo  *db.OrderMenusRepository
	orderStatusRepo *db.OrderStatusRepository
}

func NewOrderHandler(orderRepo *db.OrderRepository, orderMenusRepo *db.OrderMenusRepository, orderStatusRepo *db.OrderStatusRepository) *OrderHandler {
	return &OrderHandler{
		orderRepo:       orderRepo,
		orderMenusRepo:  orderMenusRepo,
		orderStatusRepo: orderStatusRepo,
	}
}

func (h *OrderHandler) MakeOrder(w http.ResponseWriter, r *http.Request) {

}

func (h *OrderHandler) GetOrders(w http.ResponseWriter, r *http.Request) {

}

func (h *OrderHandler) UpdateOrder(w http.ResponseWriter, r *http.Request) {

}
