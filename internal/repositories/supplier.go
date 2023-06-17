package repositories

import "hangryAPI/internal/models"

type SupplierRepositoryI interface {
	Create(*models.Supplier) error
	GetAll() ([]*models.Supplier, error)
	GetSupplierByID(int) (*models.Supplier, error)
	Update(*models.Supplier) error
	Delete(int) error
}
