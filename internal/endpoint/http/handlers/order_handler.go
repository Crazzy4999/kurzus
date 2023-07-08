package handler

import (
	"encoding/json"
	request "hangryAPI/internal/endpoint/http/requests"
	"hangryAPI/internal/endpoint/http/responses"
	"hangryAPI/internal/models"
	"hangryAPI/internal/repositories/db"
	"hangryAPI/internal/service/token"
	"hangryAPI/internal/util"
	"net/http"
)

type OrderHandler struct {
	orderRepo       *db.OrderRepository
	orderStatusRepo *db.OrderStatusRepository
}

func NewOrderHandler(orderRepo *db.OrderRepository, orderStatusRepo *db.OrderStatusRepository) *OrderHandler {
	return &OrderHandler{
		orderRepo:       orderRepo,
		orderStatusRepo: orderStatusRepo,
	}
}

func (h *OrderHandler) MakeOrder(w http.ResponseWriter, r *http.Request) {
	req := new(request.OrderRequest)
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	order := &models.Order{
		UserID:      req.UserID,
		SupplierID:  req.SupplierID,
		DriverID:    req.DriverID,
		StatusID:    req.StatusID,
		Note:        util.NullString(req.Note),
		CreatedAt:   req.CreatedAt,
		DeliveredAt: req.DeliveredAt,
	}

	err := h.orderRepo.Create(order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *OrderHandler) GetOrders(w http.ResponseWriter, r *http.Request) {
	claims, err := token.GetClaimsFromContext(r)
	if err != nil {
		http.Error(w, "invalid credentials", http.StatusUnauthorized)
		return
	}

	orders, err := h.orderRepo.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ordersResponse := &responses.OrdersCollectionResponse{}

	for _, o := range orders {
		if o.UserID == claims.ID {
			orderResponse := &responses.OrderResponse{
				ID:          o.ID,
				UserID:      o.UserID,
				SupplierID:  o.SupplierID,
				DriverID:    o.DriverID,
				StatusID:    o.StatusID,
				Note:        o.Note.String,
				CreatedAt:   o.CreatedAt,
				DeliveredAt: o.DeliveredAt,
			}

			ordersResponse.Orders = append(ordersResponse.Orders, orderResponse)
		}
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ordersResponse)
}

func (h *OrderHandler) UpdateOrder(w http.ResponseWriter, r *http.Request) {
	req := new(request.OrderRequest)
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	order := &models.Order{
		ID:          req.ID,
		StatusID:    req.StatusID,
		Note:        util.NullString(req.Note),
		DeliveredAt: req.DeliveredAt,
	}

	err := h.orderRepo.Update(order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
