package model

type Order struct {
	Items        []string `json:"items"`
	Supplier     Supplier `json:"supplier"`
	CustomerName string   `json:"customerName"`
	Price        uint     `json:"price"`
	DeliveryFee  uint     `json:"deliveryFee"`
	DeliveryTime uint     `json:"deliveryTime"`
}
