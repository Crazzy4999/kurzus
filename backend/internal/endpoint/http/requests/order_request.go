package request

type OrderRequest struct {
	ID                int    `json:"id"`
	UserID            int    `json:"userID"`
	AddressID         int    `json:"addressID"`
	SupplierID        int    `json:"supplierID"`
	DriverID          int64  `json:"driverID"`
	StatusID          int    `json:"statusID"`
	Note              string `json:"note"`
	EstimatedDelivery string `json:"estimatedDelivery"`
	DeliveredAt       string `json:"deliveredAt"`
}
