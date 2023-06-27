package responses

import "hangryAPI/internal/models"

type SupplierResponse struct {
	Id           int                 `json:"id"`
	Image        string              `json:"image"`
	Name         string              `json:"name"`
	Type         string              `json:"type"`
	WorkingHours models.WorkingHours `json:"workingHours"`
}
