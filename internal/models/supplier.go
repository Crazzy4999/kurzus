package models

const (
	RESTAURANT = iota + 1
	SUPERMARKET
	COFFEE_SHOP
)

type Supplier struct {
	ID           int          `json:"id"`
	Type         int          `json:"type"`
	Image        string       `json:"image"`
	Name         string       `json:"name"`
	Email        string       `json:"email"`
	Password     string       `json:"password"`
	DeliveryFee  float32      `json:"deliveryFee"`
	WorkingHours WorkingHours `json:"workingHours"`
}
