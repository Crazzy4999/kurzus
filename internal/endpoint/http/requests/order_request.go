package request

type OrderRequest struct {
	ID          int    `json:"id"`
	UserID      int    `json:"userID"`
	SupplierID  int    `json:"supplierID"`
	DriverID    int64  `json:"driverID"`
	StatusID    int    `json:"statusID"`
	Note        string `json:"note"`
	DeliveredAt string `json:"deliveredAt"`
}
