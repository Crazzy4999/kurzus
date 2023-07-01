package repository

import "hangryAPI/internal/models"

type AddressRepositoryI interface {
	Create(*models.Address) error
	GetAll() ([]*models.Address, error)
	GetAddressByID(int) (*models.Address, error)
	Update(*models.Address) error
	Delete(int) error
}
