package repository

import "hangryAPI/internal/models"

type DriverRepositoryI interface {
	Create(*models.Driver) error
	GetAll() ([]*models.Driver, error)
	GetAllNotDelivering() ([]*models.Driver, error)
	GetDriverByID(int) (*models.Driver, error)
	Update(*models.Driver) error
	Delete(int) error
}
