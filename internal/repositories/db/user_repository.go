package db

import (
	"database/sql"
	"hangryAPI/internal/models"
	repository "hangryAPI/internal/repositories"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) repository.UserRepositoryI {
	return &UserRepository{
		db: db,
	}
}

func (ur *UserRepository) Create(u *models.User) error {
	return nil
}

func (ur *UserRepository) GetAll() ([]*models.User, error) {
	return nil, nil
}

func (ur *UserRepository) GetUserByID(id int) (*models.User, error) {
	return nil, nil
}

func (ur *UserRepository) Update(u *models.User) error {
	return nil
}

func (ur *UserRepository) Delete(id int) error {
	return nil
}
