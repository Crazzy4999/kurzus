package responses

type OrderMenuResponse struct {
	OrderID  int `json:"orderID"`
	MenuID   int `json:"menuID"`
	Quantity int `json:"quantity"`
}
