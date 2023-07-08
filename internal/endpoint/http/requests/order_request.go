package request

type OrderRequest struct {
	ID          int    `json:"id"`
	UserID      int    `json:"userID"`
	SupplierID  int    `json:"supplierID"`
	DriverID    int    `json:"driverID"`
	StatusID    int    `json:"statusID"`
	Note        string `json:"note"`
	CreatedAt   string `json:"createdAt"`
	DeliveredAt string `json:"deliveredAt"`
}
