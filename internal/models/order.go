package model

type Order struct {
	Id           int     `json:"id"`
	UserId       int     `json:"userId"`
	SupplierId   int     `json:"supplierId"`
	DriverId     int     `json:"driverId"`
	AddressId    int     `json:"addressId"`
	StatusId     int     `json:"statusId"`
	Comment      string  `json:"comment"`
	DeliveryFee  float32 `json:"deliveryFee"`
	DeliveryTime string  `json:"deliveryTime"`
}
