package db

import "hangryAPI/internal/models"

type SupplierRepository struct{}

func NewSupplierRepository() *SupplierRepository {
	return &SupplierRepository{}
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
