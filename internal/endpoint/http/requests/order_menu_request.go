package request

type OrderMenuRequest struct {
	OrderID  int `json:"orderID"`
	MenuID   int `json:"menuID"`
	Quantity int `json:"quantity"`
}
