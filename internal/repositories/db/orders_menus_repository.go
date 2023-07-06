package db

import (
	"database/sql"
	"errors"
	"hangryAPI/internal/models"
)

type OrderMenusRepository struct {
	db *sql.DB
}

func NewOrderMenusRepository(db *sql.DB) *OrderMenusRepository {
	return &OrderMenusRepository{
		db: db,
	}
}

func (repo *OrderMenusRepository) Create(om *models.OrdersMenus) error {
	stmt, err := repo.db.Prepare("INSERT INTO orders_menus (order_id, menu_id, quantity) VALUES ($1, $2, $3)")
	if err != nil {
		return errors.New("couldn't prepare statement to create orders_menus record")
	}

	_, err = stmt.Exec(om.OrderID, om.MenuID, om.Quantity)
	if err != nil {
		return errors.New("creating orders_menus record failed")
	}

	return nil
}

func (repo *OrderMenusRepository) GetMenusByOrderID(id int) ([]*models.Menu, error) {
	stmt, err := repo.db.Prepare("SELECT order_id, menu_id, quantity FROM orders_menus WHERE orders_menus.id = $1")
	if err != nil {
		return nil, errors.New("couldn't prepare statement to get menus by order id")
	}

	rows, err := stmt.Query(id)
	if err != nil {
		return nil, errors.New("getting menus by order id failed")
	}
	defer rows.Close()

	var menuIds []int
	for rows.Next() {
		menuId := 0
		err = rows.Scan(&menuId)
		if err == sql.ErrNoRows {
			return nil, nil
		} else if err != nil {
			return nil, errors.New("types mismatch during the scanning")
		}
		menuIds = append(menuIds, menuId)
	}

	menuRepo := NewMenuRepository(repo.db)

	menus, err := menuRepo.GetAll()
	if err != nil {
		return nil, errors.New("getting menus failed")
	}

	var orderMenus []*models.Menu
	for _, menu := range menus {
		for _, id := range menuIds {
			if id == menu.ID {
				orderMenus = append(orderMenus, menu)
			}
		}
	}

	return orderMenus, nil
}

func (repo *OrderMenusRepository) DeleteOrdersByMenuID(id int) error {
	stmt, err := repo.db.Prepare("DELETE FROM orders_menus WHERE orders_menus.menu_id = $1")
	if err != nil {
		return errors.New("couldn't prepare statement to delete orders by menu id")
	}

	_, err = stmt.Exec(id)
	if err != nil {
		return errors.New("deleting orders by menu id failed")
	}

	return nil
}

func (repo *OrderMenusRepository) DeleteMenusByOrderID(id int) error {
	stmt, err := repo.db.Prepare("DELETE FROM orders_menus WHERE orders_menus.order_id = $1")
	if err != nil {
		return errors.New("couldn't prepare statement to delete menus by order id")
	}

	_, err = stmt.Exec(id)
	if err != nil {
		return errors.New("deleting menus by order id failed")
	}

	return nil
}
