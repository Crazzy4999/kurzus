package responses

type MenuResponse struct {
	ID         int     `json:"id"`
	Image      string  `json:"image"`
	SupplierID int     `json:"supplierID"`
	CategoryID int     `json:"categoryID"`
	Price      float32 `json:"price"`
}
