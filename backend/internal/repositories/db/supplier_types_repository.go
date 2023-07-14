package db

import (
	"database/sql"
	"errors"
	"hangryAPI/internal/models"
)

type SupplierTypesRepository struct {
	db *sql.DB
}

func NewSupplierTypesRepository(db *sql.DB) *SupplierTypesRepository {
	return &SupplierTypesRepository{
		db: db,
	}
}

func (repo *SupplierTypesRepository) GetTypeByID(id int) (*string, error) {
	stmt, err := repo.db.Prepare("SELECT id, type FROM supplier_types WHERE supplier_types.id = $1")
	if err != nil {
		return nil, errors.New("couldn't prepare statement to get supplier type by id")
	}

	supplierType := &models.SupplierType{}

	row := stmt.QueryRow(id)
	err = row.Scan(&supplierType.ID, &supplierType.Type)
	if err != nil {
		return nil, errors.New("getting supplier type by id failed")
	}

	return &supplierType.Type, nil
}
