package handler

import (
	"encoding/json"
	request "hangryAPI/internal/endpoint/http/requests"
	"hangryAPI/internal/endpoint/http/responses"
	"hangryAPI/internal/models"
	"hangryAPI/internal/repositories/db"
	"net/http"
	"regexp"
	"strconv"
)

type MenuHandler struct {
	menuRepo       *db.MenuRepository
	itemsMenusRepo *db.ItemsMenusRepository
}

func NewMenuHandler(menuRepo *db.MenuRepository, itemsMenusRepo *db.ItemsMenusRepository) *MenuHandler {
	return &MenuHandler{
		menuRepo:       menuRepo,
		itemsMenusRepo: itemsMenusRepo,
	}
}

func (h *MenuHandler) AddMenu(w http.ResponseWriter, r *http.Request) {
	req := new(request.MenuRequest)
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, JSON_TRANSFORM_FAILED, http.StatusBadRequest)
		return
	}

	menu := &models.Menu{
		ID:         req.ID,
		Name:       req.Name,
		Image:      req.Image,
		SupplierID: req.SupplierID,
		CategoryID: req.CategoryID,
		Price:      req.Price,
	}

	err := h.menuRepo.Create(menu)
	if err != nil {
		http.Error(w, CREATING_MENU_FAILED, http.StatusBadRequest)
		return
	}

	menus, err := h.menuRepo.GetAll()
	if err != nil {
		http.Error(w, GET_ALL_MENU_FAILED, http.StatusBadRequest)
		return
	}

	for _, m := range menus {
		saveID := m.ID
		m.ID = 0

		if *m == *menu {
			menu.ID = saveID
			break
		}
	}

	for _, i := range req.Items {
		itemMenu := &models.ItemsMenusIDPair{
			ItemID: i,
			MenuID: menu.ID,
		}

		err = h.itemsMenusRepo.Create(itemMenu)
		if err != nil {
			http.Error(w, CREATING_MENU_ITEM_FAILED, http.StatusBadRequest)
			return
		}
	}

	w.WriteHeader(http.StatusOK)
}

func (h *MenuHandler) GetMenusBySupplierID(w http.ResponseWriter, r *http.Request) {
	regex := regexp.MustCompile("\\d+")
	match := regex.FindString(r.URL.Path)

	supplierID, err := strconv.Atoi(match)
	if err != nil {
		http.Error(w, STRING_CONVERSION_FAILED, http.StatusBadRequest)
		return
	}

	menus, err := h.menuRepo.GetAll()
	if err != nil {
		http.Error(w, GET_ALL_MENU_FAILED, http.StatusBadRequest)
		return
	}

	menusResponse := &responses.MenuCollectionResponse{}

	for _, menu := range menus {
		if menu.SupplierID == supplierID {
			menuResponse := &responses.MenuResponse{
				ID:         menu.ID,
				Name:       menu.Name,
				SupplierID: menu.SupplierID,
				CategoryID: menu.CategoryID,
				Image:      menu.Image,
				Price:      menu.Price,
			}

			menusResponse.Menus = append(menusResponse.Menus, menuResponse)
		}
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(menusResponse)
}

func (h *MenuHandler) GetMenus(w http.ResponseWriter, r *http.Request) {
	menus, err := h.menuRepo.GetAll()
	if err != nil {
		http.Error(w, GET_ALL_MENU_FAILED, http.StatusBadRequest)
		return
	}

	menusResponse := &responses.MenuCollectionResponse{}

	for _, menu := range menus {
		menuResponse := &responses.MenuResponse{
			ID:         menu.ID,
			Name:       menu.Name,
			Image:      menu.Image,
			SupplierID: menu.SupplierID,
			CategoryID: menu.CategoryID,
			Price:      menu.Price,
		}

		menusResponse.Menus = append(menusResponse.Menus, menuResponse)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(menusResponse)
}

func (h *MenuHandler) RemoveMenu(w http.ResponseWriter, r *http.Request) {
	req := new(request.MenuRequest)
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, JSON_TRANSFORM_FAILED, http.StatusBadRequest)
		return
	}

	err := h.menuRepo.Delete(req.ID)
	if err != nil {
		http.Error(w, DELETING_MENU_FAILED, http.StatusBadRequest)
		return
	}

	err = h.itemsMenusRepo.DeleteByMenuID(req.ID)
	if err != nil {
		http.Error(w, DELETING_MENU_ITEM_BY_MENU_FAILED, http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
