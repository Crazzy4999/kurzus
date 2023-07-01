package models

import "database/sql"

type Address struct {
	ID          int            `json:"id"`
	City        string         `json:"city"`
	Street      string         `json:"street"`
	HouseNumber string         `json:"houseNumber"`
	ZipCode     string         `json:"zipCode"`
	FloorNumber sql.NullString `json:"floorNumber"`
	Apartment   sql.NullString `json:"apartment"`
}
