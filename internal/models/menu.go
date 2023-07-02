package models

type Menu struct {
	ID          int64  `json:"id"`
	SupplierID  int64  `json:"supplier"`
	CategorieID int64  `json:"categorie"`
	Image       string `json:"image"`
	Price       int    `json:"price"`
}
