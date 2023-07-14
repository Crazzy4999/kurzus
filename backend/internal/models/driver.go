package models

type Driver struct {
	ID           int    `json:"id"`
	IsDelivering bool   `json:"isDelivering"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	Email        string `json:"email"`
	Password     string `json:"password"`
}
