package handler

import (
	"encoding/json"
	request "hangryAPI/internal/endpoint/http/requests"
	"hangryAPI/internal/endpoint/http/responses"
	"hangryAPI/internal/models"
	"hangryAPI/internal/repositories/db"
	"hangryAPI/internal/service/token"
	"hangryAPI/internal/util"
	"net/http"
	"regexp"
	"strconv"
)

type AddressHandler struct {
	addressRepo *db.AddressRepository
}

func NewAddressHandler(addressRepo *db.AddressRepository) *AddressHandler {
	return &AddressHandler{
		addressRepo: addressRepo,
	}
}

func (h *AddressHandler) AddAddress(w http.ResponseWriter, r *http.Request) {
	req := new(request.AddressRequest)
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, JSON_TRANSFORM_FAILED, http.StatusBadRequest)
		return
	}

	claims, err := token.GetClaimsFromContext(r)
	if err != nil {
		http.Error(w, INVALID_CREDENTIALS, http.StatusBadRequest)
		return
	}

	address := &models.Address{
		UserID:      claims.ID,
		IsActive:    false,
		City:        req.City,
		Street:      req.Street,
		HouseNumber: req.HouseNumber,
		ZipCode:     req.ZipCode,
		FloorNumber: util.NullString(req.FloorNumber),
		Apartment:   util.NullString(req.Apartment),
	}

	err = h.addressRepo.Create(address)
	if err != nil {
		http.Error(w, CREATING_ADDRESS_FAILED, http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *AddressHandler) GetAddressByID(w http.ResponseWriter, r *http.Request) {
	regex := regexp.MustCompile("\\d+")
	match := regex.FindString(r.URL.Path)

	addressID, err := strconv.Atoi(match)
	if err != nil {
		http.Error(w, STRING_CONVERSION_FAILED, http.StatusBadRequest)
		return
	}

	addresses, err := h.addressRepo.GetAll()
	if err != nil {
		http.Error(w, GET_ADDRESS_BY_ID_FAILED, http.StatusBadRequest)
		return
	}

	address := &models.Address{}

	for _, a := range addresses {
		if a.ID == addressID {
			address = &models.Address{
				ID:          a.ID,
				UserID:      a.ID,
				IsActive:    a.IsActive,
				City:        a.City,
				Street:      a.Street,
				HouseNumber: a.HouseNumber,
				ZipCode:     a.ZipCode,
				FloorNumber: a.FloorNumber,
				Apartment:   a.Apartment,
			}
		}
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(address)
}

func (h *AddressHandler) GetAddressesByUserID(w http.ResponseWriter, r *http.Request) {
	claims, err := token.GetClaimsFromContext(r)
	if err != nil {
		http.Error(w, INVALID_CREDENTIALS, http.StatusUnauthorized)
		return
	}

	addresses, err := h.addressRepo.GetAll()
	if err != nil {
		http.Error(w, GET_ALL_ADDRESS_FAILED, http.StatusBadRequest)
		return
	}

	addressesResponse := &responses.AddressCollectionResponse{}
	for _, a := range addresses {
		if a.UserID == claims.ID {
			a := &responses.AddressResponse{
				ID:          a.ID,
				UserID:      a.UserID,
				IsActive:    a.IsActive,
				City:        a.City,
				Street:      a.Street,
				HouseNumber: a.HouseNumber,
				ZipCode:     a.ZipCode,
				FloorNumber: a.FloorNumber.String,
				Apartment:   a.Apartment.String,
			}

			addressesResponse.Addresses = append(addressesResponse.Addresses, a)
		}
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(addressesResponse)
}

func (h *AddressHandler) UpdateAddress(w http.ResponseWriter, r *http.Request) {
	req := new(request.AddressRequest)
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, JSON_TRANSFORM_FAILED, http.StatusBadRequest)
		return
	}

	address := &models.Address{
		ID:          req.ID,
		UserID:      req.UserID,
		IsActive:    req.IsActive,
		City:        req.City,
		Street:      req.Street,
		HouseNumber: req.HouseNumber,
		ZipCode:     req.ZipCode,
		FloorNumber: util.NullString(req.FloorNumber),
		Apartment:   util.NullString(req.Apartment),
	}

	err := h.addressRepo.Update(address)
	if err != nil {
		http.Error(w, UPDATING_ADDRESS_FAILED, http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *AddressHandler) RemoveAddress(w http.ResponseWriter, r *http.Request) {
	req := new(request.AddressRequest)
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, JSON_TRANSFORM_FAILED, http.StatusBadRequest)
		return
	}

	err := h.addressRepo.Delete(req.ID)
	if err != nil {
		http.Error(w, DELETING_ADDRESS_FAILED, http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
