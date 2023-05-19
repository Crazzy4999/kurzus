package repo

import (
	"httpA/db/model"
)

type SupplierRepository interface {
	GetAll() ([]model.Supplier, error)
}
