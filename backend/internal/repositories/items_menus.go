package repository

import "hangryAPI/internal/models"

type ItemsMenusRepositoryI interface {
	Create(*models.ItemsMenusIDPair) error
	GetAll() ([]*models.ItemsMenusIDPair, error)
	GetItemsByMenuID(int) ([]*models.Item, error)
	DeleteByItemID(int) error
	DeleteByMenuID(int) error
}
