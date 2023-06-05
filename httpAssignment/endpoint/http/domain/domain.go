package domain

import (
	"encoding/json"
	"httpA/db/repo/file"
)

func GetAllSuppliers() string {
	sr := new(file.SuppliersRepository)
	sups, _ := sr.GetAll()
	bytes, _ := json.Marshal(sups)
	return string(bytes)
}
