package db

import (
	"database/sql"
	"errors"
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
	stmt, err := sr.db.Prepare("INSERT INTO suppliers (image, name, email, password, type, opening, closing) VALUES ($1, $2, $3, $4, $5, $6, $7)")
	if err != nil {
		return errors.New("couldn't prepare statement for creating supplier")
	}

	_, err = stmt.Exec(s.Image, s.Name, s.Email, s.Password, s.Type, s.WorkingHours.Opening, s.WorkingHours.Closing)
	if err != nil {
		return errors.New("insert into suppliers failed")
	}
	return nil
}

func (sr *SupplierRepository) GetAll() ([]*models.Supplier, error) {
	stmt, err := sr.db.Prepare("SELECT id, address_id, image, name, email, password, type, opening, closing FROM suppliers")
	if err != nil {
		return nil, errors.New("couldn't prepare statement for getting all suppliers")
	}

	rows, err := stmt.Query()
	if err != nil {
		return nil, errors.New("getting all suppliers failed")
	}
	defer rows.Close()

	var suppliers []*models.Supplier
	for rows.Next() {
		supplier := &models.Supplier{}
		err = rows.Scan(&supplier.Address, &supplier.Image, &supplier.Name, &supplier.Email, &supplier.Password, &supplier.Type, &supplier.WorkingHours.Opening, &supplier.WorkingHours.Closing)
		if err == sql.ErrNoRows {
			return nil, nil
		} else if err != nil {
			return nil, errors.New("types mismatch during the scanning")
		}
		suppliers = append(suppliers, supplier)
	}

	return suppliers, nil
}

func (sr *SupplierRepository) GetSupplierByID(id int) (*models.Supplier, error) {
	stmt, err := sr.db.Prepare("SELECT id, address_id, image, name, email, password, type, opening, closing FROM suppliers WHERE suppliers.id = $1")
	if err != nil {
		return nil, errors.New("couldn't prepare statement for getting supplier by id")
	}

	supplier := &models.Supplier{}

	row := stmt.QueryRow(id)
	err = row.Scan(&supplier.ID, &supplier.Address, &supplier.Image, &supplier.Name, &supplier.Email, &supplier.Password, &supplier.Type, &supplier.WorkingHours.Opening, &supplier.WorkingHours.Closing)
	if err != nil {
		return nil, errors.New("getting supplier by id failed")
	}

	return supplier, nil
}

func (sr *SupplierRepository) Update(s *models.Supplier) error {
	stmt, err := sr.db.Prepare("UPDATE suppliers SET address_id = $1, image = $2, name = $3, type = $4, opening = $5, closing = $6 WHERE suppliers.id = $7")
	if err != nil {
		return errors.New("couldn't prepare statement for updating supplier")
	}

	_, err = stmt.Exec(s.Address, s.Image, s.Name, s.Type, s.WorkingHours.Opening, s.WorkingHours.Closing, s.ID)
	if err != nil {
		return errors.New("updating supplier failed")
	}

	return nil
}

func (sr *SupplierRepository) Delete(id int) error {
	stmt, err := sr.db.Prepare("DELETE FROM suppliers WHERE suppliers.id = $1")
	if err != nil {
		return errors.New("couldn't prepare statement for deleting supplier")
	}

	_, err = stmt.Exec(id)
	if err != nil {
		return errors.New("deleting supplier failed")
	}

	return nil
}
