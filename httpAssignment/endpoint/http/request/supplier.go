package request

import (
	"httpA/db/model"
	"httpA/endpoint/http/domain"
)

const (
	nonExistingIDErr = "There is no supplier with such ID!"
)

type UpdateSupplierRequest struct {
	Id           int                `json:"id"`
	Name         string             `json:"name"`
	Type         model.SupplierType `json:"type"`
	Image        string             `json:"image"`
	WorkingHours struct {
		Opening string `json:"opening"`
		Closing string `json:"closing"`
	}
}

func (upr *UpdateSupplierRequest) Validate() []string {
	var errors []string

	supplier := domain.FindSupplierByID(upr.Id)
	if supplier == nil {
		errors = append(errors, nonExistingIDErr)
	}

	return errors
}
