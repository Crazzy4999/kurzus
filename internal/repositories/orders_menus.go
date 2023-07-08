package repository

import "hangryAPI/internal/models"

type OrderMenusRepositoryI interface {
	Create(*models.OrdersMenus) error
	GetMenusByOrderID(int) ([]*models.Menu, error)
	Update(int) error
	DeleteOrdersByMenuID(int) error
	DeleteMenusByOrderID(int) error
}
