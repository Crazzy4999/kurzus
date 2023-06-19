package db

import (
	"database/sql"
	"hangryAPI/internal/models"
)

type DriverRepository struct {
	db *sql.DB
}

func NewDriverRepository(db *sql.DB) *DriverRepository {
	return &DriverRepository{
		db: db,
	}
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
