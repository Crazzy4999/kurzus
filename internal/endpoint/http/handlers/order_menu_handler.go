package handler

import (
	"encoding/json"
	request "hangryAPI/internal/endpoint/http/requests"
	"hangryAPI/internal/endpoint/http/responses"
	"hangryAPI/internal/models"
	"hangryAPI/internal/repositories/db"
	"net/http"
	"regexp"
	"strconv"
)

type OrderMenuHandler struct {
	orderMenuRepo *db.OrderMenusRepository
}

func NewOrderMenuHandler(orderMenuRepo *db.OrderMenusRepository) *OrderMenuHandler {
	return &OrderMenuHandler{
		orderMenuRepo: orderMenuRepo,
	}
}

func (h *OrderMenuHandler) AddOrderMenu(w http.ResponseWriter, r *http.Request) {
	req := new(request.OrderMenuRequest)
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, JSON_TRANSFORM_FAILED, http.StatusBadRequest)
		return
	}

	orderMenu := &models.OrdersMenus{
		OrderID:  req.OrderID,
		MenuID:   req.MenuID,
		Quantity: req.Quantity,
	}

	err := h.orderMenuRepo.Create(orderMenu)
	if err != nil {
		http.Error(w, CREATING_ORDER_MENU_FAILED, http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *OrderMenuHandler) GetMenusByOrderID(w http.ResponseWriter, r *http.Request) {
	regex := regexp.MustCompile("\\d+")
	match := regex.FindString(r.URL.Path)

	orderID, err := strconv.Atoi(match)
	if err != nil {
		http.Error(w, STRING_CONVERSION_FAILED, http.StatusBadRequest)
		return
	}

	menus, err := h.orderMenuRepo.GetMenusByOrderID(orderID)
	if err != nil {
		http.Error(w, GET_ALL_MENU_BY_ORDER_FAILED, http.StatusBadRequest)
		return
	}

	menusResponse := &responses.MenuCollectionResponse{}

	for _, menu := range menus {
		menuResponse := &responses.MenuResponse{
			ID:         menu.ID,
			Name:       menu.Name,
			Image:      menu.Image,
			SupplierID: menu.SupplierID,
			CategoryID: menu.CategoryID,
			Price:      menu.Price,
		}

		menusResponse.Menus = append(menusResponse.Menus, menuResponse)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(menusResponse)
}

func (h *OrderMenuHandler) GetOrderMenus(w http.ResponseWriter, r *http.Request) {
	regex := regexp.MustCompile("\\d+")
	match := regex.FindString(r.URL.Path)

	orderID, err := strconv.Atoi(match)
	if err != nil {
		http.Error(w, STRING_CONVERSION_FAILED, http.StatusBadRequest)
		return
	}

	orderMenus, err := h.orderMenuRepo.GetOrderMenusByOrderID(orderID)
	if err != nil {
		http.Error(w, GET_ALL_ORDER_MENU_BY_ORDER_ID_FAILED, http.StatusBadRequest)
		return
	}

	orderMenusCollectionResponse := &responses.OrderMenusCollectionResponse{}

	for _, orderMenu := range orderMenus {
		orderMenusResponse := &responses.OrderMenuResponse{
			OrderID:  orderMenu.OrderID,
			MenuID:   orderMenu.MenuID,
			Quantity: orderMenu.Quantity,
		}

		orderMenusCollectionResponse.OrderMenus = append(orderMenusCollectionResponse.OrderMenus, orderMenusResponse)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(orderMenusCollectionResponse)
}

func (h *OrderMenuHandler) UpdateOrderMenu(w http.ResponseWriter, r *http.Request) {
	req := new(request.OrderMenuRequest)
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, JSON_TRANSFORM_FAILED, http.StatusBadRequest)
		return
	}

	err := h.orderMenuRepo.UpdateByOrderID(req.OrderID)
	if err != nil {
		http.Error(w, UPDATING_ORDER_MENU_BY_ORDER_FAILED, http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *OrderMenuHandler) RemoveOrderMenu(w http.ResponseWriter, r *http.Request) {
	req := new(request.OrderMenuRequest)
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, JSON_TRANSFORM_FAILED, http.StatusBadRequest)
		return
	}

	err := h.orderMenuRepo.DeleteByOrderID(req.OrderID)
	if err != nil {
		http.Error(w, DELETING_ORDER_MENU_BY_ORDER_FAILED, http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
