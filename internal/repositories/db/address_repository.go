package db

import (
	"database/sql"
	"errors"
	"hangryAPI/internal/models"
)

type AddressRepository struct {
	db *sql.DB
}

func NewAddressRepository(db *sql.DB) *AddressRepository {
	return &AddressRepository{
		db: db,
	}
}

func (ar *AddressRepository) Create(a *models.Address) error {
	stmt, err := ar.db.Prepare("INSERT INTO addresses (city, street, house_number, zip_code, floor_number, apartment) VALUES ($1, $2, $3, $4, $5, $6)")
	if err != nil {
		return errors.New("couldn't prepare statement to create address")
	}

	_, err = stmt.Exec(a.City, a.Street, a.HouseNumber, a.ZipCode, a.FloorNumber, a.Apartment)
	if err != nil {
		return errors.New("failed to crate address")
	}

	return nil
}

func (ar *AddressRepository) GetAll() ([]*models.Address, error) {
	stmt, err := ar.db.Prepare("SELECT id, city, street, house_number, zip_code, floor_number, apartment FROM addresses")
	if err != nil {
		return nil, errors.New("couldn't prepare statement to create address")
	}

	rows, err := stmt.Query()
	if err != nil {
		return nil, errors.New("getting all addresses failed")
	}
	defer rows.Close()

	var addresses []*models.Address
	for rows.Next() {
		address := &models.Address{}
		err = rows.Scan(&address.ID, &address.City, &address.Street, &address.HouseNumber, &address.ZipCode, &address.FloorNumber, &address.Apartment)
		if err == sql.ErrNoRows {
			return nil, nil
		} else if err != nil {
			return nil, errors.New("types mismatch during the scanning")
		}
		addresses = append(addresses, address)
	}

	return addresses, nil
}

func (ar *AddressRepository) GetAddressByID(id int) (*models.Address, error) {
	stmt, err := ar.db.Prepare("SELECT id, city, street, house_number, zip_code, floor_number, apartment FROM addresses WHERE addresses.id = $1")
	if err != nil {
		return nil, errors.New("couldn't prepare statement to create address")
	}

	address := &models.Address{}
	row := stmt.QueryRow(id)
	err = row.Scan(&address.ID, &address.City, &address.Street, &address.HouseNumber, &address.ZipCode, &address.FloorNumber, &address.Apartment)
	if err != nil {
		return nil, errors.New("getting address by id failed")
	}

	return address, nil
}

func (ar *AddressRepository) Update(a *models.Address) error {
	stmt, err := ar.db.Prepare("UPDATE addresses SET city = $1, street = $2, house_number = $3, zip_code = $4, floor_number = $5, apartment = $6 WHERE addresses.id = $7")
	if err != nil {
		return errors.New("couldn't prepare statement to update address")
	}

	_, err = stmt.Exec(a.City, a.Street, a.HouseNumber, a.ZipCode, a.FloorNumber, a.Apartment, a.ID)
	if err != nil {
		return errors.New("updating address failed")
	}

	return nil
}

func (ar *AddressRepository) Delete(id int) error {
	stmt, err := ar.db.Prepare("DELETE FROM addresses WHERE addresses.id = $1")
	if err != nil {
		return errors.New("couldn't prepare statement to delete address")
	}

	_, err = stmt.Exec(id)
	if err != nil {
		return errors.New("deleting address failed")
	}

	return nil
}
