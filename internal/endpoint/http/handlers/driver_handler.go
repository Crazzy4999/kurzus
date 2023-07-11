package handler

import (
	"encoding/json"
	request "hangryAPI/internal/endpoint/http/requests"
	"hangryAPI/internal/endpoint/http/responses"
	"hangryAPI/internal/models"
	"hangryAPI/internal/repositories/db"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type DriverHandler struct {
	driverRepo *db.DriverRepository
}

func NewDriverHandler(driverRepo *db.DriverRepository) *DriverHandler {
	return &DriverHandler{
		driverRepo: driverRepo,
	}
}

func (h *DriverHandler) AddDriver(w http.ResponseWriter, r *http.Request) {
	req := new(request.DriverRequest)
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, JSON_TRANSFORM_FAILED, http.StatusBadRequest)
		return
	}

	password, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, PASSWORD_GENERATION_FAILED, http.StatusBadRequest)
		return
	}

	driver := &models.Driver{
		ID:        req.ID,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Password:  string(password),
	}

	err = h.driverRepo.Create(driver)
	if err != nil {
		http.Error(w, CREATING_DRIVER_FAILED, http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *DriverHandler) GetDrivers(w http.ResponseWriter, r *http.Request) {
	drivers, err := h.driverRepo.GetAll()
	if err != nil {
		http.Error(w, GET_ALL_DRIVER_FAILED, http.StatusBadRequest)
		return
	}

	driversResponse := &responses.DriverCollectionResponse{}

	for _, driver := range drivers {
		driverResponse := &responses.DriverResponse{
			ID:        driver.ID,
			FirstName: driver.FirstName,
			LastName:  driver.LastName,
			Email:     driver.Email,
			Password:  driver.Password,
		}

		driversResponse.Drivers = append(driversResponse.Drivers, driverResponse)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(driversResponse)
}

func (h *DriverHandler) UpdateDriver(w http.ResponseWriter, r *http.Request) {
	req := new(request.DriverRequest)
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, JSON_TRANSFORM_FAILED, http.StatusBadRequest)
		return
	}

	driver := &models.Driver{
		ID:        req.ID,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Password:  req.Password,
	}

	err := h.driverRepo.Update(driver)
	if err != nil {
		http.Error(w, UPDATING_DRIVER_FAILED, http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *DriverHandler) RemoveDriver(w http.ResponseWriter, r *http.Request) {
	req := new(request.DriverRequest)
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, JSON_TRANSFORM_FAILED, http.StatusBadRequest)
		return
	}

	err := h.driverRepo.Delete(req.ID)
	if err != nil {
		http.Error(w, DELETING_DRIVER_FAILED, http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
