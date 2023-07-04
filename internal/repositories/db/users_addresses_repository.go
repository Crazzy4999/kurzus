package db

import (
	"database/sql"
	"errors"
	"hangryAPI/internal/models"
)

type UsersAddressesRepository struct {
	db *sql.DB
}

func NewUsersAddressesRepository(db *sql.DB) *UsersAddressesRepository {
	return &UsersAddressesRepository{
		db: db,
	}
}

func (repo *UsersAddressesRepository) Create(ua *models.UserAddressIDPair) error {
	stmt, err := repo.db.Prepare("INSERT INTO users_addresses (user_id, address_id) VALUES ($1, $2)")
	if err != nil {
		return errors.New("couldn't prepare statement to create users_addresses record")
	}

	_, err = stmt.Exec(ua.UserID, ua.AddressID)
	if err != nil {
		return errors.New("creating users_addresses record failed")
	}

	return nil
}

func (repo *UsersAddressesRepository) GetAddressesForUserByID(id int) ([]*models.Address, error) {
	stmt, err := repo.db.Prepare("SELECT address_id FROM users_addresses WHERE users_addresses.user_id = $1")
	if err != nil {
		return nil, errors.New("couldn't prepare statement to get addresses")
	}

	rows, err := stmt.Query(id)
	if err != nil {
		return nil, errors.New("getting addresses failed")
	}
	defer rows.Close()

	addressRepo := NewAddressRepository(repo.db)

	addresses, err := addressRepo.GetAll()
	if err != nil {
		return nil, errors.New("getting addresses failed")
	}

	var address_ids []*int
	for rows.Next() {
		address_id := 0

		err = rows.Scan(&address_id)
		if err == sql.ErrNoRows {
			return nil, nil
		} else if err != nil {
			return nil, errors.New("types mismatch during the scanning")
		}

		if address_id == 0 {
			return nil, errors.New("scan failed on address")
		}
		address_ids = append(address_ids, &address_id)
	}

	var userAddresses []*models.Address
	for _, address := range addresses {
		for _, address_id := range address_ids {
			if address_id == &address.ID {
				userAddresses = append(userAddresses, address)
			}
		}
	}

	return userAddresses, nil
}

func (repo *UsersAddressesRepository) DeleteByUserID(id int) error {
	stmt, err := repo.db.Prepare("DELETE FROM users_addresses WHERE users_address.user_id = $1")
	if err != nil {
		return errors.New("couldn't prepare statement to delete address/es by user id")
	}

	_, err = stmt.Exec(id)
	if err != nil {
		return errors.New("deleting address/es by user id failed")
	}

	return nil
}

func (repo *UsersAddressesRepository) DeleteByAddressID(id int) error {
	stmt, err := repo.db.Prepare("DELETE FROM users_addresses WHERE users_address.address_id = $1")
	if err != nil {
		return errors.New("couldn't prepare statement to delete address by address id")
	}

	_, err = stmt.Exec(id)
	if err != nil {
		return errors.New("deleting address by address id failed")
	}

	return nil
}
