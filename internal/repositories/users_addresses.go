package repository

import "hangryAPI/internal/models"

type UsersAddressesRepositoryI interface {
	Create(*models.UsersAddressesIDPair) error
	GetAddressesForUserByID(int) ([]*models.Address, error)
	DeleteByUserID(int) error
	DeleteByAddressID(int) error
}
