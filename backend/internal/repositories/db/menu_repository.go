package db

import (
	"database/sql"
	"errors"
	"hangryAPI/internal/models"
)

type MenuRepository struct {
	db *sql.DB
}

func NewMenuRepository(db *sql.DB) *MenuRepository {
	return &MenuRepository{
		db: db,
	}
}

func (repo *MenuRepository) Create(m *models.Menu) error {
	stmt, err := repo.db.Prepare("INSERT INTO menus (name, image, supplier_id, category_id, price) VALUES ($1, $2, $3, $4, $5)")
	if err != nil {
		return errors.New("couldn't prepare statement to create menu")
	}

	_, err = stmt.Exec(m.Name, m.Image, m.SupplierID, m.CategoryID, m.Price)
	if err != nil {
		return errors.New("failed to create menu")
	}

	return nil
}

func (repo *MenuRepository) GetAll() ([]*models.Menu, error) {
	stmt, err := repo.db.Prepare("SELECT id, name, supplier_id, category_id, image, price FROM menus")
	if err != nil {
		return nil, errors.New("couldn't prepare statement for getting all menus")
	}

	rows, err := stmt.Query()
	if err != nil {
		return nil, errors.New("getting all menus failed")
	}
	defer rows.Close()

	var menus []*models.Menu
	for rows.Next() {
		menu := &models.Menu{}
		err = rows.Scan(&menu.ID, &menu.Name, &menu.SupplierID, &menu.CategoryID, &menu.Image, &menu.Price)
		if err == sql.ErrNoRows {
			return nil, nil
		} else if err != nil {
			return nil, errors.New("types mismatch during the scanning")
		}
		menus = append(menus, menu)
	}

	return menus, nil
}

func (repo *MenuRepository) GetMenuByID(id int) (*models.Menu, error) {
	stmt, err := repo.db.Prepare("SELECT id, name, supplier_id, category_id, image, price FROM menus WHERE menus.id = $1")
	if err != nil {
		return nil, errors.New("couldn't prepare statement for getting all menus")
	}

	menu := &models.Menu{}

	row := stmt.QueryRow(id)
	err = row.Scan(&menu.ID, &menu.Name, &menu.SupplierID, &menu.CategoryID, &menu.Image, &menu.Price)
	if err != nil {
		return nil, errors.New("getting menu by id failed")
	}

	return menu, nil
}

func (repo *MenuRepository) Update(m *models.Menu) error {
	stmt, err := repo.db.Prepare("UPDATE menus SET name = $1, image = $2, price = $3 WHERE menus.id = $4")
	if err != nil {
		return errors.New("couldn't prepare statement for updating menu")
	}

	_, err = stmt.Exec(m.Name, m.Image, m.Price, m.ID)
	if err != nil {
		return errors.New("updating menu failed")
	}

	return nil
}

func (repo *MenuRepository) Delete(id int) error {
	stmt, err := repo.db.Prepare("DELETE FROM menus WHERE menus.id = $1")
	if err != nil {
		return errors.New("couldn't prepare statement to delete menu")
	}

	_, err = stmt.Exec(id)
	if err != nil {
		return errors.New("deleting menu failed")
	}

	return nil
}
