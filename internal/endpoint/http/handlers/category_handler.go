package handler

import (
	"encoding/json"
	request "hangryAPI/internal/endpoint/http/requests"
	"hangryAPI/internal/endpoint/http/responses"
	"hangryAPI/internal/models"
	"hangryAPI/internal/repositories/db"
	"net/http"
)

type CategoryHandler struct {
	categoriesRepo *db.CategoriesRepository
}

func NewCategoryHandler(categoriesRepo *db.CategoriesRepository) *CategoryHandler {
	return &CategoryHandler{
		categoriesRepo: categoriesRepo,
	}
}

func (h *CategoryHandler) AddCategory(w http.ResponseWriter, r *http.Request) {
	req := new(request.CategoryRequest)
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	category := &models.Category{
		ID:   req.ID,
		Name: req.Name,
	}

	err := h.categoriesRepo.Create(category)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *CategoryHandler) GetCategories(w http.ResponseWriter, r *http.Request) {
	categories, err := h.categoriesRepo.GetCategorieAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	categoriesResponse := &responses.CategoryCollectionResponse{}

	for _, c := range categories {
		categoryResponse := &responses.CategoryResponse{
			ID:   c.ID,
			Name: c.Name,
		}

		categoriesResponse.Categories = append(categoriesResponse.Categories, categoryResponse)
	}

	w.WriteHeader(http.StatusOK)
}

func (h *CategoryHandler) RemoveCategory(w http.ResponseWriter, r *http.Request) {
	req := new(request.CategoryRequest)
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := h.categoriesRepo.Delete(req.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
