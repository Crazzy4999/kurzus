package request

import (
	"fmt"
	"httpA/db/model"
	"httpA/db/repo/file"
)

const (
	maxAddressLength = 64
	addressToolong   = "Address excedes maximum number of characters allowed: 64!"
	noSuchSupplier   = "No such supplier with this ID!"
)

type CreateOrderRequest struct {
	OrderId    int `json:"order_id"`
	SupplierId int `json:"supplier_id"`
	UserId     int `json:"user_id"`
	Address    struct {
		Street   string `json:"street"`
		Building string `json:"building"`
		Apt      int    `json:"apt"`
	}
	Status model.StatusType `json:"status"`
}

func (cro *CreateOrderRequest) Validate() []string {
	var errors []string

	supr := &file.SuppliersRepository{}

	sups, err := supr.GetAll()
	if err != nil {
		fmt.Println(err)
	}
	supplierID := &sups[cro.SupplierId-1]

	if supplierID == nil {
		errors = append(errors, noSuchSupplier)
	}

	if maxAddressLength < len(cro.Address.Street) {
		errors = append(errors, addressToolong)
	}

	return errors
}
