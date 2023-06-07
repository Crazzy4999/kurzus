package domain

import (
	"httpA/db/model"
	"httpA/db/repo/file"
)

func GetAllSuppliers() []model.Supplier {
	sr := new(file.SuppliersRepository)
	sups, _ := sr.GetAll()
	return sups
}

func FindSupplierByID(id int) *model.Supplier {
	sr := new(file.SuppliersRepository)
	sups, _ := sr.GetAll()
	return &sups[id]
}

func getSupplierIndex(supplier model.Supplier) int {
	for i, s := range GetAllSuppliers() {
		if s.Id == supplier.Id {
			return i
		}
	}
	return -1
}

func UpdateSupplier(supplier *model.Supplier) {
	supplierIndex := getSupplierIndex(*supplier)
	GetAllSuppliers()[supplierIndex] = *supplier
}
