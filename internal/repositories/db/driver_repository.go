package db

import "hangryAPI/internal/models"

type DriverRepository struct{}

func NewDriverRepository() *DriverRepository {
	return &DriverRepository{}
}

func (dr *DriverRepository) Create(*models.Driver) error {
	return nil
}

func (dr *DriverRepository) GetAll() ([]*models.Driver, error) {
	return nil, nil
}

func (dr *DriverRepository) GetDriverByID(int) (*models.Driver, error) {
	return nil, nil
}

func (dr *DriverRepository) Update(*models.Driver) error {
	return nil
}

func (dr *DriverRepository) Delete(int) error {
	return nil
}
