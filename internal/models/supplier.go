package models

const (
	RESTAURANT  = "restaurant"
	SUPERMARKET = "supermarket"
	COFFEE_SHOP = "coffee_shop"
)

type Supplier struct {
	ID           int          `json:"id"`
	Name         string       `json:"name"`
	Email        string       `json:"email"`
	Password     string       `json:"password"`
	Type         string       `json:"type"`
	WorkingHours WorkingHours `json:"workingHours"`
	Address      string       `json:"address"`
}
