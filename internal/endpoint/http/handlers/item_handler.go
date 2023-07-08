package handler

import (
	"encoding/json"
	request "hangryAPI/internal/endpoint/http/requests"
	"hangryAPI/internal/endpoint/http/responses"
	"hangryAPI/internal/models"
	"hangryAPI/internal/repositories/db"
	"net/http"
)

type ItemHandler struct {
	itemRepo       *db.ItemRepository
	itemsMenusRepo *db.ItemsMenusRepository
}

func NewItemHandler(itemRepo *db.ItemRepository, itemsMenusRepo *db.ItemsMenusRepository) *ItemHandler {
	return &ItemHandler{
		itemRepo:       itemRepo,
		itemsMenusRepo: itemsMenusRepo,
	}
}

func (h *ItemHandler) AddItem(w http.ResponseWriter, r *http.Request) {
	req := new(request.ItemRequest)
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	item := &models.Item{
		Ingerdient: req.Ingerdient,
	}

	err := h.itemRepo.Create(item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *ItemHandler) GetItems(w http.ResponseWriter, r *http.Request) {
	items, err := h.itemRepo.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	itemsResponse := &responses.ItemCollectionResponse{}

	for _, i := range items {
		itemResponse := &responses.ItemResponse{
			ID:         i.ID,
			Ingerdient: i.Ingerdient,
		}

		itemsResponse.Items = append(itemsResponse.Items, itemResponse)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(itemsResponse)
}

func (h *ItemHandler) RemoveItem(w http.ResponseWriter, r *http.Request) {
	req := new(request.ItemRequest)
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := h.itemRepo.Delete(req.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.itemsMenusRepo.DeleteByItemID(req.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
