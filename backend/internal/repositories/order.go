package repository

import "hangryAPI/internal/models"

type OrderRepositoryI interface {
	Create(*models.Order) error
	GetAll() ([]*models.Order, error)
	GetOrdersWithCreatedStatus() ([]*models.Order, error)
	GetOrdersWithDeliveringStatus() ([]*models.Order, error)
	GetOrderByID(int) (*models.Order, error)
	Update(*models.Order) error
	Delete(int) error
}
