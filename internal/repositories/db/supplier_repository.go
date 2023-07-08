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

func (repo *SupplierRepository) Create(s *models.Supplier) error {
	stmt, err := repo.db.Prepare("INSERT INTO suppliers (type, image, name, email, password, delivery_fee, opening, closing) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)")
	if err != nil {
		return errors.New("couldn't prepare statement for creating supplier")
	}

	_, err = stmt.Exec(s.Type, s.Image, s.Name, s.Email, s.Password, s.DeliveryFee, s.WorkingHours.Opening, s.WorkingHours.Closing)
	if err != nil {
		return errors.New("insert into suppliers failed")
	}
	return nil
}

func (repo *SupplierRepository) GetAll() ([]*models.Supplier, error) {
	stmt, err := repo.db.Prepare("SELECT id, type, image, name, email, password, delivery_fee, opening, closing FROM suppliers")
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
		err = rows.Scan(&supplier.Type, &supplier.Image, &supplier.Name, &supplier.Email, &supplier.Password, &supplier.DeliveryFee, &supplier.WorkingHours.Opening, &supplier.WorkingHours.Closing)
		if err == sql.ErrNoRows {
			return nil, nil
		} else if err != nil {
			return nil, errors.New("types mismatch during the scanning")
		}
		suppliers = append(suppliers, supplier)
	}

	return suppliers, nil
}

func (repo *SupplierRepository) GetSupplierByID(id int) (*models.Supplier, error) {
	stmt, err := repo.db.Prepare("SELECT id, type, image, name, email, password, delivery_fee, opening, closing FROM suppliers WHERE suppliers.id = $1")
	if err != nil {
		return nil, errors.New("couldn't prepare statement for getting supplier by id")
	}

	supplier := &models.Supplier{}

	row := stmt.QueryRow(id)
	err = row.Scan(&supplier.ID, &supplier.Type, &supplier.Image, &supplier.Name, &supplier.Email, &supplier.Password, &supplier.DeliveryFee, &supplier.WorkingHours.Opening, &supplier.WorkingHours.Closing)
	if err != nil {
		return nil, errors.New("getting supplier by id failed")
	}

	return supplier, nil
}

func (repo *SupplierRepository) Update(s *models.Supplier) error {
	stmt, err := repo.db.Prepare("UPDATE suppliers SET type = $1, image = $2, name = $3, delivery_fee = $4, opening = $5, closing = $6 WHERE suppliers.id = $7")
	if err != nil {
		return errors.New("couldn't prepare statement for updating supplier")
	}

	_, err = stmt.Exec(s.Type, s.Image, s.Name, s.DeliveryFee, s.WorkingHours.Opening, s.WorkingHours.Closing, s.ID)
	if err != nil {
		return errors.New("updating supplier failed")
	}

	return nil
}

func (repo *SupplierRepository) Delete(id int) error {
	stmt, err := repo.db.Prepare("DELETE FROM suppliers WHERE suppliers.id = $1")
	if err != nil {
		return errors.New("couldn't prepare statement for deleting supplier")
	}

	_, err = stmt.Exec(id)
	if err != nil {
		return errors.New("deleting supplier failed")
	}

	return nil
}
