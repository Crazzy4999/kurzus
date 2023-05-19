package repo

import (
	"httpA/db/model"
)

type OrdersRepository interface {
	Create(*model.Order) (model.Order, error)
}
