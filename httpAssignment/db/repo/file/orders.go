package file

import (
	"bufio"
	"encoding/json"
	"fmt"
	"httpA/db/model"
	"io"
	"io/ioutil"
	"os"
)

type OrdersRepository struct{}

func (o *OrdersRepository) GetAll() ([]model.Order, error) {
	orders := make([]model.Order, 0)
	info, err := ioutil.ReadDir("data/orders")
	if err != nil {
		fmt.Println(err)
	}
	for _, file := range info {
		order := new(model.Order)
		content, err := os.Open(fmt.Sprintf("data/orders/%s", file.Name()))
		if err != nil {
			fmt.Println(err)
		}
		defer content.Close()
		bytes, err := io.ReadAll(content)
		_ = json.Unmarshal(bytes, order)
		if err != nil {
			fmt.Println(err)
		}
		orders = append(orders, *order)
	}
	return orders, nil
}

func (o *OrdersRepository) OrdersCount() int {
	orders, err := o.GetAll()
	if err != nil {
		fmt.Println(err)
	}
	return len(orders) + 1
}

func (o *OrdersRepository) Create(order *model.Order) (model.Order, error) {
	ordersCount := o.OrdersCount()
	fileName := fmt.Sprintf("data/orders/%d.json", ordersCount)

	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_EXCL|os.O_RDWR, 0755)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	orderBytes, err := json.Marshal(*order)
	if err != nil {
		fmt.Println(err)
	}

	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(string(orderBytes))
	if err != nil {
		fmt.Println(err)
	}
	defer writer.Flush()

	return *order, nil
}

func (o *OrdersRepository) GetOrderByID(id int) model.Order {
	orders, err := o.GetAll()
	if err != nil {
		fmt.Println(err)
	}
	return orders[id]
}

func (o *OrdersRepository) DeleteOrder(id int) {
	orderPath := fmt.Sprintf("data/orders/%d.json", id)
	err := os.Remove(orderPath)
	if err != nil {
		fmt.Println(err)
	}
}
