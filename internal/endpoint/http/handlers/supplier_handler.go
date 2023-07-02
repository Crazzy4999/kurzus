package handler

import (
	config "hangryAPI/configs"
	"hangryAPI/internal/repositories/db"
	"net/http"
)

type SupplierHandler struct {
	supplierRepo *db.SupplierRepository
	cfg          *config.Config
}

func NewSupplierHandler(supplierRepo *db.SupplierRepository, cfg *config.Config) *SupplierHandler {
	return &SupplierHandler{
		supplierRepo: supplierRepo,
		cfg:          cfg,
	}
}

func (h *SupplierHandler) AddSupplier(w http.ResponseWriter, r *http.Request) {

}
