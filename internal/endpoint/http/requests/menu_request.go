package request

type MenuRequest struct {
	ID         int     `json:"id"`
	Items      []int   `json:"items"`
	Image      string  `json:"image"`
	SupplierID int     `json:"supplierID"`
	CategoryID int     `json:"categoryID"`
	Price      float32 `json:"price"`
}
