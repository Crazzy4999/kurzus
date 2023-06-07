package domain

import (
	"httpA/db/model"
	"httpA/db/repo/file"
)

func AddOrder(order *model.Order) {
	or := &file.OrdersRepository{}
	or.Create(order)
}

func DeleteOrder(id int) {
	or := &file.OrdersRepository{}
	or.DeleteOrder(id)
}
