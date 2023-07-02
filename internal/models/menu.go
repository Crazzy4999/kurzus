package models

type Menu struct {
	ID          int64  `json:"id"`
	Image       string `json:"image"`
	SupplierID  int64  `json:"supplier"`
	CategorieID int64  `json:"categorie"`
	Quantity    int    `json:"quantity"`
	Price       int    `json:"price"`
}
