package models

import "database/sql"

type Order struct {
	ID          int            `json:"id"`
	UserID      int            `json:"userID"`
	SupplierID  int            `json:"supplierID"`
	DriverID    int            `json:"driverID"`
	StatusID    int            `json:"statusID"`
	Note        sql.NullString `json:"note"`
	DeliveryFee float32        `json:"deliveryFee"`
	CreatedAt   string         `json:"createdAt"`
	DeliveredAt string         `json:"deliveredAt"`
}
