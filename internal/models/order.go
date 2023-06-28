package models

type Order struct {
	ID           int     `json:"id"`
	UserID       int     `json:"userID"`
	SupplierID   int     `json:"supplierID"`
	DriverID     int     `json:"driverID"`
	AddressID    int     `json:"addressID"`
	StatusID     int     `json:"statusID"`
	Comment      string  `json:"comment"`
	DeliveryFee  float32 `json:"deliveryFee"`
	DeliveryTime string  `json:"deliveryTime"`
}
