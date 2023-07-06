package db

import (
	"database/sql"
	"errors"
	"fmt"
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

func (repo *AddressRepository) Create(a *models.Address) error {
	stmt, err := repo.db.Prepare("INSERT INTO addresses (user_id, is_active, city, street, house_number, zip_code, floor_number, apartment) VALUES ($1, $2, $3, $4, $5, $6, $7)")
	if err != nil {
		return errors.New("couldn't prepare statement to create address")
	}

	_, err = stmt.Exec(a.UserID, false, a.City, a.Street, a.HouseNumber, a.ZipCode, a.FloorNumber, a.Apartment)
	if err != nil {
		return errors.New("failed to crate address")
	}

	return nil
}

func (repo *AddressRepository) GetAll() ([]*models.Address, error) {
	stmt, err := repo.db.Prepare("SELECT id, user_id, is_active, city, street, house_number, zip_code, floor_number, apartment FROM addresses")
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
		err = rows.Scan(&address.ID, &address.UserID, &address.IsActive, &address.City, &address.Street, &address.HouseNumber, &address.ZipCode, &address.FloorNumber, &address.Apartment)
		if err == sql.ErrNoRows {
			return nil, nil
		} else if err != nil {
			return nil, errors.New("types mismatch during the scanning")
		}
		addresses = append(addresses, address)
	}

	return addresses, nil
}

func (repo *AddressRepository) GetAddressByID(id int) (*models.Address, error) {
	stmt, err := repo.db.Prepare("SELECT id, user_id, is_active, city, street, house_number, zip_code, floor_number, apartment FROM addresses WHERE addresses.id = $1")
	if err != nil {
		return nil, errors.New("couldn't prepare statement to create address")
	}

	address := &models.Address{}
	row := stmt.QueryRow(id)
	err = row.Scan(&address.ID, &address.UserID, &address.IsActive, &address.City, &address.Street, &address.HouseNumber, &address.ZipCode, &address.FloorNumber, &address.Apartment)
	if err != nil {
		return nil, errors.New("getting address by id failed")
	}

	return address, nil
}

func (repo *AddressRepository) Update(a *models.Address) error {
	stmt, err := repo.db.Prepare("UPDATE addresses SET isActive = $1, city = $2, street = $3, house_number = $4, zip_code = $5, floor_number = $6, apartment = $7 WHERE addresses.id = $8")
	if err != nil {
		return errors.New("couldn't prepare statement to update address")
	}

	_, err = stmt.Exec(a.IsActive, a.City, a.Street, a.HouseNumber, a.ZipCode, a.FloorNumber, a.Apartment, a.ID)
	if err != nil {
		return errors.New("updating address failed")
	}

	return nil
}

func (repo *AddressRepository) Delete(id int64) error {
	stmt, err := repo.db.Prepare("DELETE FROM addresses WHERE addresses.id = $1")
	if err != nil {
		return errors.New("couldn't prepare statement to delete address")
	}

	_, err = stmt.Exec(id)
	if err != nil {
		fmt.Println(err)
		return errors.New("deleting address failed")
	}

	return nil
}
