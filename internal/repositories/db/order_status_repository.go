package db

import (
	"database/sql"
	"errors"
	"hangryAPI/internal/models"
)

type OrderStatusRepository struct {
	db *sql.DB
}

func NewOrderStatusRepository(db *sql.DB) *OrderStatusRepository {
	return &OrderStatusRepository{
		db: db,
	}
}

func (repo *OrderRepository) GetStatusByID(id int) (*string, error) {
	stmt, err := repo.db.Prepare("SELECT id, status FROM order_status WHERE order_status.id = $1")
	if err != nil {
		return nil, errors.New("couldn't prepare statement to get order status by id")
	}

	orderStatus := &models.OrderStatus{}

	row := stmt.QueryRow(id)
	err = row.Scan(&orderStatus.ID, &orderStatus.Status)
	if err != nil {
		return nil, errors.New("getting order status by id failed")
	}

	return &orderStatus.Status, nil
}
