package repository

import "hangryAPI/internal/models"

type UserRepositoryI interface {
	Create(*models.User) error
	GetAll() ([]*models.User, error)
	GetUserByEmail(string) (*models.User, error)
	Update(*models.User) error
	Delete(int) error
}
