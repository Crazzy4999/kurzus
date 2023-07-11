package handler

import (
	"encoding/json"
	"hangryAPI/internal/endpoint/http/responses"
	"hangryAPI/internal/repositories/db"
	"net/http"
	"regexp"
	"strconv"
)

type ItemsMenusHandler struct {
	itemsMenusRepo *db.ItemsMenusRepository
	menuRepo       *db.MenuRepository
}

func NewItemsMenusHandler(itemsMenusRepo *db.ItemsMenusRepository, menuRepo *db.MenuRepository) *ItemsMenusHandler {
	return &ItemsMenusHandler{
		itemsMenusRepo: itemsMenusRepo,
		menuRepo:       menuRepo,
	}
}

func (h *ItemsMenusHandler) GetItemsByMenuID(w http.ResponseWriter, r *http.Request) {
	regex := regexp.MustCompile("\\d+")
	match := regex.FindString(r.URL.Path)

	menuID, err := strconv.Atoi(match)
	if err != nil {
		http.Error(w, STRING_CONVERSION_FAILED, http.StatusBadRequest)
		return
	}

	items, err := h.itemsMenusRepo.GetItemsByMenuID(menuID)
	if err != nil {
		http.Error(w, GET_ALL_MENU_FAILED, http.StatusBadRequest)
	}

	itemsResponse := &responses.ItemCollectionResponse{}

	for _, item := range items {
		itemResponse := &responses.ItemResponse{
			ID:         item.ID,
			Ingerdient: item.Ingerdient,
		}

		itemsResponse.Items = append(itemsResponse.Items, itemResponse)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(itemsResponse)
}
