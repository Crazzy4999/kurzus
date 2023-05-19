package model

type SupplierType string

const (
	RESTAURANT  SupplierType = "restaurant"
	COFFEE_SHOP SupplierType = "coffee_shop"
	SUPERMARKET SupplierType = "supermarket"
)

type Supplier struct {
	Id      int          `json:"id"`
	Name    string       `json:"name"`
	Type    SupplierType `json:"type"`
	Image   string       `json:"image"`
	Opening string       `json:"opening"`
	Closing string       `json:"closing"`
}