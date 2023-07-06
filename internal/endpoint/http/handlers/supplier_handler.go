package handler

import (
	"encoding/json"
	request "hangryAPI/internal/endpoint/http/requests"
	"hangryAPI/internal/endpoint/http/responses"
	"hangryAPI/internal/models"
	"hangryAPI/internal/repositories/db"
	"net/http"
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
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	supplier := &models.Supplier{
		Type:         req.Type,
		Image:        req.Image,
		Name:         req.Name,
		Email:        req.Email,
		Password:     req.Password,
		WorkingHours: req.WorkingHours,
	}

	suppliers, err := h.supplierRepo.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusOK)
		return
	}

	for _, s := range suppliers {
		if s.Email == supplier.Email {
			http.Error(w, "supplier already exists", http.StatusBadRequest)
			return
		}
	}

	err = h.supplierRepo.Create(supplier)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *SupplierHandler) GetSuppliers(w http.ResponseWriter, r *http.Request) {
	suppliers, err := h.supplierRepo.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	suppliersResponse := &responses.SupplierCollectionResponse{}

	for _, s := range suppliers {
		supplierType, err := h.supplierTypesRepo.GetTypeByID(s.Type)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		supplierResponse := &responses.SupplierResponse{
			ID:           s.ID,
			Type:         *supplierType,
			Image:        s.Image,
			Name:         s.Name,
			WorkingHours: s.WorkingHours,
		}

		suppliersResponse.Suppliers = append(suppliersResponse.Suppliers, *supplierResponse)
	}

	w.WriteHeader(http.StatusOK)
}

func (h *SupplierHandler) UpdateSupplier(w http.ResponseWriter, r *http.Request) {
	req := new(request.SupplierRequest)
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	supplier := &models.Supplier{
		ID:           req.ID,
		Type:         req.Type,
		Image:        req.Image,
		Name:         req.Name,
		Email:        req.Email,
		Password:     req.Password,
		WorkingHours: req.WorkingHours,
	}

	err := h.supplierRepo.Update(supplier)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *SupplierHandler) RemoveSupplier(w http.ResponseWriter, r *http.Request) {
	req := new(request.SupplierRequest)
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := h.supplierRepo.Delete(req.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
