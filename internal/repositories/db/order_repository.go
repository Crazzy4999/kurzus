package db

import (
	"database/sql"
	"errors"
	"fmt"
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

func (repo *OrderRepository) Create(o *models.Order) error {
	stmt, err := repo.db.Prepare("INSERT INTO orders (user_id, address_id, supplier_id, driver_id, status_id, note, estimated_delivery) VALUES ($1, $2, $3, $4, $5, $6, $7)")
	if err != nil {
		return errors.New("couldn't prepare statenemt to create order")
	}

	_, err = stmt.Exec(o.UserID, o.AddressID, o.SupplierID, o.DriverID, models.CREATED, o.Note, o.EstimatedDelivery)
	if err != nil {
		return errors.New("insert into orders failed")
	}

	return nil
}

func (or *OrderRepository) GetAll() ([]*models.Order, error) {
	stmt, err := or.db.Prepare("SELECT id, user_id, address_id, supplier_id, driver_id, status_id, note, created_at, delivered_at FROM orders")
	if err != nil {
		return nil, errors.New("couldn't prepare statement to getting all orders")
	}

	rows, err := stmt.Query()
	if err != nil {
		return nil, errors.New("getting all orders failed")
	}
	defer rows.Close()

	var orders []*models.Order
	for rows.Next() {
		order := &models.Order{}
		err = rows.Scan(&order.ID, &order.UserID, &order.AddressID, &order.SupplierID, &order.DriverID, &order.StatusID, &order.Note, &order.CreatedAt, &order.EstimatedDelivery, &order.DeliveredAt)
		if err == sql.ErrNoRows {
			return nil, nil
		} else if err != nil {
			fmt.Println(err)

			return nil, errors.New("types mismatch during the scanning")
		}
		orders = append(orders, order)
	}

	return orders, nil
}

func (or *OrderRepository) GetOrdersWithCreatedStatus() ([]*models.Order, error) {
	stmt, err := or.db.Prepare("SELECT id, user_id, address_id, supplier_id, driver_id, status_id, note, created_at, estimated_delivery, delivered_at FROM orders WHERE status_id = $1")
	if err != nil {
		return nil, errors.New("couldn't prepare statement to getting all orders")
	}

	rows, err := stmt.Query(models.CREATED)
	if err != nil {
		return nil, errors.New("getting all orders failed")
	}
	defer rows.Close()

	var orders []*models.Order
	for rows.Next() {
		order := &models.Order{}
		err = rows.Scan(&order.ID, &order.UserID, &order.AddressID, &order.SupplierID, &order.DriverID, &order.StatusID, &order.Note, &order.CreatedAt, &order.EstimatedDelivery, &order.DeliveredAt)
		if err == sql.ErrNoRows {
			return nil, nil
		} else if err != nil {
			fmt.Println(err)

			return nil, errors.New("types mismatch during the scanning")
		}
		orders = append(orders, order)
	}

	return orders, nil
}

func (or *OrderRepository) GetOrdersWithDeliveringStatus() ([]*models.Order, error) {
	stmt, err := or.db.Prepare("SELECT id, user_id, address_id, supplier_id, driver_id, status_id, note, created_at, estimated_delivery, delivered_at FROM orders WHERE status_id = $1")
	if err != nil {
		return nil, errors.New("couldn't prepare statement to getting all orders")
	}

	rows, err := stmt.Query(models.DELIVERING)
	if err != nil {
		return nil, errors.New("getting all orders failed")
	}
	defer rows.Close()

	var orders []*models.Order
	for rows.Next() {
		order := &models.Order{}
		err = rows.Scan(&order.ID, &order.UserID, &order.AddressID, &order.SupplierID, &order.DriverID, &order.StatusID, &order.Note, &order.CreatedAt, &order.EstimatedDelivery, &order.DeliveredAt)
		if err == sql.ErrNoRows {
			return nil, nil
		} else if err != nil {
			fmt.Println(err)

			return nil, errors.New("types mismatch during the scanning")
		}
		orders = append(orders, order)
	}

	return orders, nil
}

func (or *OrderRepository) GetOrderByID(id int) (*models.Order, error) {
	stmt, err := or.db.Prepare("SELECT id, user_id, address_id, supplier_id, driver_id, status_id, note, created_at, estimated_delivery, delivered_at FROM orders WHERE orders.id = $1")
	if err != nil {
		return nil, errors.New("couldn't prepare statement for getting order by id")
	}

	order := &models.Order{}

	row := stmt.QueryRow(id)
	err = row.Scan(&order.ID, &order.UserID, &order.AddressID, &order.SupplierID, &order.DriverID, &order.StatusID, &order.Note, &order.CreatedAt, &order.EstimatedDelivery, &order.DeliveredAt)
	if err != nil {
		return nil, errors.New("getting order by id failed")
	}

	return order, nil
}

func (or *OrderRepository) Update(o *models.Order) error {
	stmt, err := or.db.Prepare("UPDATE orders SET driver_id = $1, status_id = $2, note = $3, estimated_delivery = $4, delivered_at = $5 WHERE orders.id = $6")
	if err != nil {
		return errors.New("couldn't prepare statement for updating order")
	}

	_, err = stmt.Exec(o.DriverID, o.StatusID, o.Note, o.EstimatedDelivery, o.DeliveredAt, o.ID)
	if err != nil {
		return errors.New("updating order failed")
	}

	return nil
}

func (or *OrderRepository) Delete(id int) error {
	stmt, err := or.db.Prepare("DELETE FROM orders WHERE orders.id = $1")
	if err != nil {
		return errors.New("couldn't prepare statement for deleting order")
	}

	_, err = stmt.Exec(id)
	if err != nil {
		return errors.New("deleting order failed")
	}

	return nil
}
