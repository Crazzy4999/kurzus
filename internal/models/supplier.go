package models

import "database/sql"

const (
	RESTAURANT = iota + 1
	SUPERMARKET
	COFFEE_SHOP
)

type Supplier struct {
	ID           int           `json:"id"`
	Address      sql.NullInt64 `json:"address"`
	Type         int           `json:"type"`
	Image        string        `json:"image"`
	Name         string        `json:"name"`
	Email        string        `json:"email"`
	Password     string        `json:"password"`
	WorkingHours WorkingHours  `json:"workingHours"`
}
