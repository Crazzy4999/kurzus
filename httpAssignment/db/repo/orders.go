package repo

import (
	"httpA/db/model"
)

type OrdersRepository interface {
	GetAll() ([]model.Order, error)
	OrdersCount() int
	Create(*model.Order) (model.Order, error)
	GetOrderByID(id int) model.Order
	DeleteOrder(*model.Order)
}
