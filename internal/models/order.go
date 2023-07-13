package models

import "database/sql"

type Order struct {
	ID          int            `json:"id"`
	UserID      int            `json:"userID"`
	AddressID   int            `json:"addressID"`
	SupplierID  int            `json:"supplierID"`
	DriverID    sql.NullInt64  `json:"driverID"`
	StatusID    int            `json:"statusID"`
	Note        sql.NullString `json:"note"`
	CreatedAt   string         `json:"createdAt"`
	DeliveredAt string         `json:"deliveredAt"`
}
