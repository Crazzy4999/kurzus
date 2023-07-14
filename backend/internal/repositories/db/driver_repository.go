package db

import (
	"database/sql"
	"errors"
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

func (repo *DriverRepository) Create(d *models.Driver) error {
	stmt, err := repo.db.Prepare("INSERT INTO drivers (is_delivering, first_name, last_name, email, password) VALUES ($1, $2, $3, $4, $5)")
	if err != nil {
		return errors.New("couldn't prepare statement for creating driver")
	}

	_, err = stmt.Exec(d.IsDelivering, d.FirstName, d.LastName, d.Email, d.Password)
	if err != nil {
		return errors.New("creating driver failed")
	}

	return nil
}

func (repo *DriverRepository) GetAll() ([]*models.Driver, error) {
	stmt, err := repo.db.Prepare("SELECT id, is_delivering, first_name, last_name, email, password FROM drivers")
	if err != nil {
		return nil, errors.New("couldn't prepare statement for getting all drivers")
	}

	rows, err := stmt.Query()
	if err != nil {
		return nil, errors.New("getting all drivers failed")
	}
	defer rows.Close()

	var drivers []*models.Driver
	for rows.Next() {
		driver := &models.Driver{}
		err = rows.Scan(&driver.ID, &driver.IsDelivering, &driver.FirstName, &driver.LastName, &driver.Email, &driver.Password)
		if err == sql.ErrNoRows {
			return nil, nil
		} else if err != nil {
			return nil, errors.New("types mismatch during the scanning")
		}
		drivers = append(drivers, driver)
	}
	return drivers, nil
}

func (repo *DriverRepository) GetAllNotDelivering() ([]*models.Driver, error) {
	stmt, err := repo.db.Prepare("SELECT id, is_delivering, first_name, last_name, email, password FROM drivers WHERE is_delivering = $1")
	if err != nil {
		return nil, errors.New("couldn't prepare statement for getting all drivers")
	}

	rows, err := stmt.Query(false)
	if err != nil {
		return nil, errors.New("getting all drivers failed")
	}
	defer rows.Close()

	var drivers []*models.Driver
	for rows.Next() {
		driver := &models.Driver{}
		err = rows.Scan(&driver.ID, &driver.IsDelivering, &driver.FirstName, &driver.LastName, &driver.Email, &driver.Password)
		if err == sql.ErrNoRows {
			return nil, nil
		} else if err != nil {
			return nil, errors.New("types mismatch during the scanning")
		}
		drivers = append(drivers, driver)
	}

	return drivers, nil
}

func (repo *DriverRepository) GetDriverByID(id int) (*models.Driver, error) {
	stmt, err := repo.db.Prepare("SELECT id, is_delivering, first_name, last_name, email, password FROM drivers WHERE drivers.id = $1")
	if err != nil {
		return nil, errors.New("couldn't prepare statement for getting all drivers")
	}

	driver := &models.Driver{}
	row := stmt.QueryRow(id)
	err = row.Scan(&driver.ID, &driver.IsDelivering, &driver.FirstName, &driver.LastName, &driver.Email, &driver.Password)
	if err != nil {
		return nil, errors.New("getting driver by id failed")
	}

	return driver, nil
}

func (repo *DriverRepository) Update(d *models.Driver) error {
	stmt, err := repo.db.Prepare("UPDATE drivers SET is_delivering = $1, first_name = $2, last_name = $3, password = $4 WHERE drivers.id = $5")
	if err != nil {
		return errors.New("couldn't prepare statement for updating driver")
	}

	_, err = stmt.Exec(d.IsDelivering, d.FirstName, d.LastName, d.Password, d.ID)
	if err != nil {
		return errors.New("updating driver failed")
	}

	return nil
}

func (repo *DriverRepository) Delete(id int) error {
	stmt, err := repo.db.Prepare("DELETE FROM drivers WHERE drivers.id = $1")
	if err != nil {
		return errors.New("couldn't prepare statement for deleting driver")
	}

	_, err = stmt.Exec(id)
	if err != nil {
		return errors.New("deleting driver failed")
	}

	return nil
}
