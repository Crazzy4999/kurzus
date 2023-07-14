package models

type Menu struct {
	ID         int     `json:"id"`
	Name       string  `json:"name"`
	Image      string  `json:"image"`
	SupplierID int     `json:"supplierID"`
	CategoryID int     `json:"categoryID"`
	Price      float32 `json:"price"`
}
