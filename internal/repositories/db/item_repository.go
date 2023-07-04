package db

import (
	"database/sql"
	"errors"
	"hangryAPI/internal/models"
)

type ItemRepository struct {
	db *sql.DB
}

func NewItemRepository(db *sql.DB) *ItemRepository {
	return &ItemRepository{
		db: db,
	}
}

func (repo *ItemRepository) Create(m *models.Item) error {
	stmt, err := repo.db.Prepare("INSERT INTO items (menu_id, ingredient) VALUES ($1, $2)")
	if err != nil {
		return errors.New("couldn't prepare statment to create item")
	}

	_, err = stmt.Exec(m.MenuId, m.Ingerdient)
	if err != nil {
		return errors.New("creating item failed")
	}

	return nil
}

func (repo *ItemRepository) GetAll() ([]*models.Item, error) {
	stmt, err := repo.db.Prepare("SELECT id, menu_id, ingredient FROM items")
	if err != nil {
		return nil, errors.New("couldn't prepare statement for getting all items")
	}

	rows, err := stmt.Query()
	if err != nil {
		return nil, errors.New("getting all items failed")
	}
	defer rows.Close()

	var items []*models.Item
	for rows.Next() {
		item := &models.Item{}
		err = rows.Scan(&item.ID, &item.MenuId, &item.Ingerdient)
		if err == sql.ErrNoRows {
			return nil, nil
		} else if err != nil {
			return nil, errors.New("types mismatch during the scanning")
		}
		items = append(items, item)
	}

	return items, nil
}

func (repo *ItemRepository) GetItemByID(id int) (*models.Item, error) {
	stmt, err := repo.db.Prepare("SELECT id, menu_id, ingredient FROM items WHERE items.id = $1")
	if err != nil {
		return nil, errors.New("couldn't prepare statement for getting item by id")
	}

	item := &models.Item{}

	row := stmt.QueryRow(id)
	err = row.Scan(&item.ID, &item.MenuId, &item.Ingerdient)
	if err != nil {
		return nil, errors.New("getting item by id failed")
	}

	return item, nil
}

func (repo *ItemRepository) Update(m *models.Item) error {
	stmt, err := repo.db.Prepare("UPDATE items SET menu_id = $1, ingredient = $2 WHERE items.id = $3")
	if err != nil {
		return errors.New("couldn't prepare statement to update item")
	}

	_, err = stmt.Exec(m.MenuId, m.Ingerdient, m.ID)
	if err != nil {
		return errors.New("updating item failed")
	}

	return nil
}

func (repo *ItemRepository) Delete(id int) error {
	stmt, err := repo.db.Prepare("DELETE FROM items WHERE items.id = $1")
	if err != nil {
		return errors.New("couldn't prepare statement to delete item")
	}

	_, err = stmt.Exec(id)
	if err != nil {
		return errors.New("deleting item failed")
	}

	return nil
}
