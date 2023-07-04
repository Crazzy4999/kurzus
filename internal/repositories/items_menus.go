package repository

import "hangryAPI/internal/models"

type ItemsMenusRepositoryI interface {
	Create(*models.ItemsMenus) error
	GetItemsByMenuID(int) ([]*models.Item, error)
	DeleteByItemID(int) error
	DeleteByMenuID(int) error
}
