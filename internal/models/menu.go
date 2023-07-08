package models

type Menu struct {
	ID         int     `json:"id"`
	SupplierID int     `json:"supplierID"`
	CategoryID int     `json:"categoryID"`
	Image      string  `json:"image"`
	Price      float32 `json:"price"`
}
