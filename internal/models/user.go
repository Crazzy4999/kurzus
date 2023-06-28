package models

import "database/sql"

type User struct {
	ID        int            `json:"id"`
	Address   sql.NullString `json:"address"`
	FirstName string         `json:"firstName"`
	LastName  string         `json:"lastName"`
	Email     string         `json:"email"`
	Password  string         `json:"password"`
}
