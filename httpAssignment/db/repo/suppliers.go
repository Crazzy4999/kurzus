package repo

import (
	"encoding/json"
	"fmt"
	"httpA/db/model"
	"io"
	"io/ioutil"
	"os"
)

type SupplierRepository interface {
	GetAll() ([]model.Supplier, error)
}

type SupplierManager struct{}

func (s *SupplierManager) GetAll() ([]model.Supplier, error) {
	suppliers := make([]model.Supplier, 0)
	info, err := ioutil.ReadDir("data/suppliers")
	if err != nil {
		fmt.Println(err)
	}
	for _, file := range info {
		supplier := new(model.Supplier)
		content, err := os.Open(fmt.Sprintf("data/suppliers/%s", file.Name()))
		if err != nil {
			fmt.Println(err)
		}
		bytes, err := io.ReadAll(content)
		_ = json.Unmarshal(bytes, supplier)
		if err != nil {
			fmt.Println(err)
		}
		suppliers = append(suppliers, *supplier)
	}
	return suppliers, err
}
