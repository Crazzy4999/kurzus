package db

import (
	"database/sql"
	"errors"
	"hangryAPI/internal/models"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (ur *UserRepository) Create(u *models.User) error {
	stmt, err := ur.db.Prepare("INSERT INTO users (first_name, last_name, email, password) VALUES ($1, $2, $3, $4);")
	if err != nil {
		return errors.New("couldn't prepare statment to insert new user")
	}

	_, err = stmt.Exec(u.FirstName, u.LastName, u.Email, u.Password)
	if err != nil {
		return errors.New("insert into users failed")
	}

	return nil
}

func (ur *UserRepository) GetAll() ([]*models.User, error) {
	return nil, nil
}

func (ur *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	return nil, nil
}

func (ur *UserRepository) Update(u *models.User) error {
	return nil
}

func (ur *UserRepository) Delete(id int) error {
	return nil
}
