package request

import (
	"hangryAPI/internal/models"
)

type SupplierRequest struct {
	ID           int                 `json:"id"`
	Type         int                 `json:"type"`
	Image        string              `json:"image"`
	Name         string              `json:"name"`
	Email        string              `json:"email"`
	Password     string              `json:"password"`
	Description  string              `json:"description"`
	DeliveryTime int                 `json:"deliveryTime"`
	DeliveryFee  float32             `json:"deliveryFee"`
	WorkingHours models.WorkingHours `json:"workingHours"`
}
