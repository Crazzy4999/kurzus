package file

import (
	"encoding/json"
	"httpA/db/model"
)

type OrdersRepository struct{}

func (o *OrdersRepository) Create(order *model.Order) (model.Order, error) {
	err := json.Unmarshal(*new([]byte), new(string))
	return *order, err
}
