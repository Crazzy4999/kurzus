package repo

import (
	"httpA/db/model"
)

type SuppliersRepository interface {
	GetAll() ([]model.Supplier, error)
}
