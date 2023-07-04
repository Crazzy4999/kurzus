package db

import (
	"database/sql"
	"errors"
	"hangryAPI/internal/models"
)

type CategoriesRepository struct {
	db *sql.DB
}

func NewCategoriesRepository(db *sql.DB) *CategoriesRepository {
	return &CategoriesRepository{
		db: db,
	}
}

func (repo *CategoriesRepository) Create(c *models.Category) error {
	stmt, err := repo.db.Prepare("INSERT INTO categories (name) VALUES ($1)")
	if err != nil {
		return errors.New("couldn't preprae statement to create category")
	}

	_, err = stmt.Exec(c.Name)
	if err != nil {
		return errors.New("creating category failed")
	}

	return nil
}

func (repo *CategoriesRepository) GetCategorieByID(id int) (*models.Category, error) {
	stmt, err := repo.db.Prepare("SELECT id, name FROM categories WHERE categories.id = $1")
	if err != nil {
		return nil, errors.New("couldn't prepare statement to get category by id")
	}

	category := &models.Category{}

	row := stmt.QueryRow(id)
	err = row.Scan(&category.ID, &category.Name)
	if err != nil {
		return nil, errors.New("getting category by id failed")
	}

	return category, nil
}

func (repo *CategoriesRepository) Delete(id int) error {
	stmt, err := repo.db.Prepare("DELETE FROM categories WHERE categories.id = $1")
	if err != nil {
		return errors.New("couldn't prepare statement to delete category")
	}

	_, err = stmt.Exec(id)
	if err != nil {
		return errors.New("deleting category failed")
	}

	return nil
}
