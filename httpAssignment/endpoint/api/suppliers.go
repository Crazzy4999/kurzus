package api

import (
	"encoding/json"
	"fmt"
	"httpA/db/model"
	"httpA/endpoint/http/domain"
	"httpA/endpoint/http/request"
	"httpA/endpoint/http/response"
	"net/http"
	"regexp"
	"strconv"
)

func GetAllSuppliers(w http.ResponseWriter, r *http.Request) {
	response.SendOK(w, domain.GetAllSuppliers())
}

func GetSupplier(w http.ResponseWriter, r *http.Request) {
	regex := regexp.MustCompile("\\d+")
	match := regex.FindString(r.URL.Path)

	id, err := strconv.Atoi(match)
	if err != nil {
		response.SendBadRequestError(w, err)
		return
	}

	supplier := domain.FindSupplierByID(id)
	if supplier == nil {
		response.SendBadRequestError(w, fmt.Errorf("no supplier with id: %d", id))
		return
	}
	response.SendOK(w, supplier)
}

func UpdateSupplier(w http.ResponseWriter, r *http.Request) {
	supr := request.UpdateSupplierRequest{}

	err := request.ParseJsonRequest(r, &supr)
	if err != nil {
		response.SendBadRequestError(w, err)
		return
	}

	regex := regexp.MustCompile("\\d+")
	id := regex.FindString(r.URL.Path)
	intId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err)
	}
	supplier := model.Supplier{Id: intId, Name: supr.Name, Type: supr.Type, Image: supr.Image, WorkingHours: struct {
		Opening string `json:"opening"`
		Closing string `json:"closing"`
	}{Opening: supr.WorkingHours.Opening, Closing: supr.WorkingHours.Closing}}
	domain.UpdateSupplier(&supplier)

	response.SendOK(w, supplier)
}

func GetAllSuppliersRefresh(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("https://foodapi.golang.nixdev.co/suppliers")
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	respBytes, err := json.Marshal(resp)
	fmt.Println(string(respBytes))
}
