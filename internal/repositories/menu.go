package repository

import "hangryAPI/internal/models"

type MenuRepositoryI interface {
	Create(*models.Menu) error
	GetAll() ([]*models.Menu, error)
	GetMenuByID(int) (*models.Menu, error)
	Update(*models.Menu) error
	Delete(int) error
}
