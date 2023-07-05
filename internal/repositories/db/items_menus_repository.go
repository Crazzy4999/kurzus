package db

import (
	"database/sql"
	"errors"
	"hangryAPI/internal/models"
)

type ItemsMenusRepository struct {
	db *sql.DB
}

func NewItemsMenusRepository(db *sql.DB) *ItemsMenusRepository {
	return &ItemsMenusRepository{
		db: db,
	}
}

func (repo *ItemsMenusRepository) Create(im *models.ItemsMenusIDPair) error {
	stmt, err := repo.db.Prepare("INSERT INTO items_menus (item_id, menu_id) VALUES ($1, $2)")
	if err != nil {
		return errors.New("couldn't prepare statement to create items_menus record")
	}

	_, err = stmt.Exec(im.ItemID, im.MenuID)
	if err != nil {
		return errors.New("creating items_menus record failed")
	}

	return nil
}

func (repo *ItemsMenusRepository) GetItemsByMenuID(id int) ([]*models.Item, error) {
	stmt, err := repo.db.Prepare("")
	if err != nil {
		return nil, errors.New("")
	}

	rows, err := stmt.Query(id)
	if err != nil {
		return nil, errors.New("")
	}

	var itemIds []*int
	for rows.Next() {
		itemId := 0
		err = rows.Scan(&itemId)
		if err == sql.ErrNoRows {
			return nil, nil
		} else if err != nil {
			return nil, errors.New("types mismatch during the scanning")
		}
		itemIds = append(itemIds, &itemId)
	}

	itemRepo := NewItemRepository(repo.db)

	items, err := itemRepo.GetAll()
	if err != nil {
		return nil, errors.New("getting all items failed")
	}

	var menuItems []*models.Item
	for _, item := range items {
		for _, id := range itemIds {
			if id == &item.ID {
				menuItems = append(menuItems, item)
			}
		}
	}

	return menuItems, nil
}

func (repo *ItemsMenusRepository) DeleteByItemID(id int) error {
	stmt, err := repo.db.Prepare("DELETE FROM items_menus WHERE items_menus.item_id = $1")
	if err != nil {
		return errors.New("couldn't prepare statement to delete items_menus record by item id")
	}

	_, err = stmt.Exec(id)
	if err != nil {
		return errors.New("deleting by item id failed")
	}

	return nil
}

func (repo *ItemsMenusRepository) DeleteByMenuID(id int) error {
	stmt, err := repo.db.Prepare("DELETE FROM items_menus WHERE items_menus.menu_id = $1")
	if err != nil {
		return errors.New("couldn't prepare statement to delete items_menus record by menu id")
	}

	_, err = stmt.Exec(id)
	if err != nil {
		return errors.New("deleting by menu id failed")
	}

	return nil
}
