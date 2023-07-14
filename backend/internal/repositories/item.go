package repository

import "hangryAPI/internal/models"

type ItemRepositoryI interface {
	Create(*models.Item) error
	GetAll() ([]*models.Item, error)
	GetItemByID(int) (*models.Item, error)
	Update(*models.Item) error
	Delete(int) error
}
