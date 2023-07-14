package responses

import "hangryAPI/internal/models"

type SupplierResponse struct {
	ID           int                 `json:"id"`
	Type         string              `json:"type"`
	Image        string              `json:"image"`
	Name         string              `json:"name"`
	Description  string              `json:"description"`
	DeliveryTime int                 `json:"deliveryTime"`
	DeliveryFee  float32             `json:"deliveryFee"`
	WorkingHours models.WorkingHours `json:"workingHours"`
}
