package responses

import "hangryAPI/internal/models"

type SupplierResponse struct {
	ID           int                 `json:"id"`
	Type         string              `json:"type"`
	Image        string              `json:"image"`
	Name         string              `json:"name"`
	DeliveryFee  float32             `json:"deliveryFee"`
	WorkingHours models.WorkingHours `json:"workingHours"`
}
