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

	"golang.org/x/crypto/bcrypt"
)

type SupplierHandler struct {
	supplierRepo      *db.SupplierRepository
	supplierTypesRepo *db.SupplierTypesRepository
}

func NewSupplierHandler(supplierRepo *db.SupplierRepository, supplierTypesRepo *db.SupplierTypesRepository) *SupplierHandler {
	return &SupplierHandler{
		supplierRepo:      supplierRepo,
		supplierTypesRepo: supplierTypesRepo,
	}
}

func (h *SupplierHandler) AddSupplier(w http.ResponseWriter, r *http.Request) {
	req := new(request.SupplierRequest)
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, JSON_TRANSFORM_FAILED, http.StatusBadRequest)
		return
	}

	password, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, PASSWORD_GENERATION_FAILED, http.StatusBadRequest)
		return
	}

	supplier := &models.Supplier{
		Type:         req.Type,
		Image:        req.Image,
		Name:         req.Name,
		Email:        req.Email,
		Password:     string(password),
		Description:  req.Description,
		DeliveryTime: req.DeliveryTime,
		DeliveryFee:  req.DeliveryFee,
		WorkingHours: req.WorkingHours,
	}

	suppliers, err := h.supplierRepo.GetAll()
	if err != nil {
		http.Error(w, GET_ALL_SUPPLIER_FAILED, http.StatusOK)
		return
	}

	for _, s := range suppliers {
		if s.Email == supplier.Email {
			http.Error(w, SUPPLIER_EXISTS, http.StatusBadRequest)
			return
		}
	}

	err = h.supplierRepo.Create(supplier)
	if err != nil {
		http.Error(w, CREATING_SUPPLIER_FAILED, http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *SupplierHandler) GetSupplierByID(w http.ResponseWriter, r *http.Request) {
	regex := regexp.MustCompile("\\d+")
	match := regex.FindString(r.URL.Path)

	supplierID, err := strconv.Atoi(match)
	if err != nil {
		http.Error(w, STRING_CONVERSION_FAILED, http.StatusBadRequest)
		return
	}

	suppliers, err := h.supplierRepo.GetAll()
	if err != nil {
		http.Error(w, GET_ALL_SUPPLIER_FAILED, http.StatusBadRequest)
		return
	}

	var supplier *responses.SupplierResponse

	for _, s := range suppliers {
		if s.ID == supplierID {
			supplierType, err := h.supplierTypesRepo.GetTypeByID(s.Type)
			if err != nil {
				http.Error(w, GET_SUPPLIER_TYPE_FAILED, http.StatusBadRequest)
				return
			}

			supplier = &responses.SupplierResponse{
				ID:           s.ID,
				Type:         *supplierType,
				Image:        s.Image,
				Name:         s.Name,
				Description:  s.Description,
				DeliveryTime: s.DeliveryTime,
				DeliveryFee:  s.DeliveryFee,
				WorkingHours: s.WorkingHours,
			}

			break
		}
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(supplier)
}

func (h *SupplierHandler) GetSuppliers(w http.ResponseWriter, r *http.Request) {
	suppliers, err := h.supplierRepo.GetAll()
	if err != nil {
		http.Error(w, GET_ALL_SUPPLIER_FAILED, http.StatusBadRequest)
		return
	}

	suppliersResponse := &responses.SupplierCollectionResponse{}

	for _, s := range suppliers {
		supplierType, err := h.supplierTypesRepo.GetTypeByID(s.Type)
		if err != nil {
			http.Error(w, GET_SUPPLIER_TYPE_FAILED, http.StatusBadRequest)
			return
		}

		supplierResponse := &responses.SupplierResponse{
			ID:           s.ID,
			Type:         *supplierType,
			Image:        s.Image,
			Name:         s.Name,
			Description:  s.Description,
			DeliveryTime: s.DeliveryTime,
			DeliveryFee:  s.DeliveryFee,
			WorkingHours: s.WorkingHours,
		}

		suppliersResponse.Suppliers = append(suppliersResponse.Suppliers, *supplierResponse)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(suppliersResponse)
}

func (h *SupplierHandler) UpdateSupplier(w http.ResponseWriter, r *http.Request) {
	req := new(request.SupplierRequest)
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, JSON_TRANSFORM_FAILED, http.StatusBadRequest)
		return
	}

	supplier := &models.Supplier{
		ID:           req.ID,
		Type:         req.Type,
		Image:        req.Image,
		Name:         req.Name,
		Email:        req.Email,
		Password:     req.Password,
		Description:  req.Description,
		DeliveryTime: req.DeliveryTime,
		DeliveryFee:  req.DeliveryFee,
		WorkingHours: req.WorkingHours,
	}

	err := h.supplierRepo.Update(supplier)
	if err != nil {
		http.Error(w, UPDATING_SUPPLIER_FAILED, http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *SupplierHandler) RemoveSupplier(w http.ResponseWriter, r *http.Request) {
	req := new(request.SupplierRequest)
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, JSON_TRANSFORM_FAILED, http.StatusBadRequest)
		return
	}

	err := h.supplierRepo.Delete(req.ID)
	if err != nil {
		http.Error(w, DELETING_SUPPLIER_FAILED, http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
