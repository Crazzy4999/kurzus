package db

import (
	"database/sql"
	"hangryAPI/internal/models"
)

type SupplierRepository struct {
	db *sql.DB
}

func NewSupplierRepository(db *sql.DB) *SupplierRepository {
	return &SupplierRepository{
		db: db,
	}
}

func (sr *SupplierRepository) Create(s *models.Supplier) error {
	return nil
}

func (sr *SupplierRepository) GetAll() ([]*models.Supplier, error) {
	return nil, nil
}

func (sr *SupplierRepository) GetSupplierByID(id int) (*models.Supplier, error) {
	return nil, nil
}

func (sr *SupplierRepository) Update(s *models.Supplier) error {
	return nil
}

func (sr *SupplierRepository) Delete(id int) error {
	return nil
}
