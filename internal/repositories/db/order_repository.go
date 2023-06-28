package db

import (
	"database/sql"
	"hangryAPI/internal/models"
)

type OrderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{
		db: db,
	}
}

func (or *OrderRepository) Create(o *models.Order) error {
	return nil
}

func (or *OrderRepository) GetAll() ([]*models.Order, error) {
	return nil, nil
}

func (or *OrderRepository) GetOrderByID(id int) (*models.Order, error) {
	return nil, nil
}

func (or *OrderRepository) Update(o *models.Order) error {
	return nil
}

func (or *OrderRepository) Delete(id int) error {
	return nil
}
