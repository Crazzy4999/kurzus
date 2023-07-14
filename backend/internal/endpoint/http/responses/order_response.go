package responses

type OrderResponse struct {
	ID                int    `json:"id"`
	UserID            int    `json:"userID"`
	AddressID         int    `json:"addressID"`
	SupplierID        int    `json:"supplierID"`
	DriverID          int64  `json:"driverID"`
	StatusID          int    `json:"statusID"`
	Note              string `json:"note"`
	CreatedAt         string `json:"createdAt"`
	EstimatedDelivery string `json:"estimatedDelivery"`
	DeliveredAt       string `json:"deliveredAt"`
}
