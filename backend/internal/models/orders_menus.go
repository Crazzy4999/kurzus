package models

type OrdersMenus struct {
	OrderID  int `json:"orderId"`
	MenuID   int `json:"menuId"`
	Quantity int `json:"quantity"`
}
