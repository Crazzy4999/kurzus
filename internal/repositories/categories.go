package repository

import "hangryAPI/internal/models"

type CategoriesRepositoryI interface {
	Create(*models.Category) error
	GetCategorieAll() ([]*models.Category, error)
	GetCategorieByID(int) (*models.Category, error)
	Delete(int) error
}
