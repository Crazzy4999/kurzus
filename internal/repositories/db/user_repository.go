package db

import (
	"hangryAPI/internal/models"
)

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
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
