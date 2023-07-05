package models

type Menu struct {
	ID          int    `json:"id"`
	SupplierID  int    `json:"supplier"`
	CategorieID int    `json:"categorie"`
	Image       string `json:"image"`
	Price       int    `json:"price"`
}
