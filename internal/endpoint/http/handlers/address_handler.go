package handler

import (
	"encoding/json"
	config "hangryAPI/configs"
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
	addressRepo        *db.AddressRepository
	usersAddressesRepo *db.UsersAddressesRepository
	userRepo           *db.UserRepository
	cfg                *config.Config
}

func NewAddressHandler(addressRepo *db.AddressRepository, usersAddressesRepo *db.UsersAddressesRepository, userRepo *db.UserRepository, cfg *config.Config) *AddressHandler {
	return &AddressHandler{
		addressRepo:        addressRepo,
		usersAddressesRepo: usersAddressesRepo,
		userRepo:           userRepo,
		cfg:                cfg,
	}
}

func (h *AddressHandler) AddAddress(w http.ResponseWriter, r *http.Request) {
	req := new(request.AddressRequest)
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	addresses, err := h.addressRepo.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	address := &models.Address{
		City:        req.City,
		Street:      req.Street,
		HouseNumber: req.HouseNumber,
		ZipCode:     req.ZipCode,
		FloorNumber: util.NullString(req.FloorNumber),
		Apartment:   util.NullString(req.Apartment),
	}

	for _, a := range addresses {
		a.ID = 0
		if *address == *a {
			w.WriteHeader(http.StatusOK)
			return
		}
	}

	err = h.addressRepo.Create(address)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *AddressHandler) GetAddressByID(w http.ResponseWriter, r *http.Request) {
	pattern := regexp.MustCompile("\\d+")
	match := pattern.FindString(r.URL.Path)
	addressID, err := strconv.Atoi(match)
	if err != nil {
		http.Error(w, "getting address id failed", http.StatusBadRequest)
		return
	}

	address, err := h.addressRepo.GetAddressByID(addressID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	addressResponse := &responses.AddressResponse{
		ID:          address.ID,
		City:        address.City,
		Street:      address.Street,
		HouseNumber: address.HouseNumber,
		ZipCode:     address.ZipCode,
		FloorNumber: address.FloorNumber.String,
		Apartment:   address.Apartment.String,
	}

	json.NewEncoder(w).Encode(addressResponse)
	w.WriteHeader(http.StatusOK)
}

func (h *AddressHandler) AddUserAddress(w http.ResponseWriter, r *http.Request) {
	req := new(request.AddressRequest)
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	address := &models.Address{
		City:        req.City,
		Street:      req.Street,
		HouseNumber: req.HouseNumber,
		ZipCode:     req.ZipCode,
		FloorNumber: util.NullString(req.FloorNumber),
		Apartment:   util.NullString(req.Apartment),
	}

	err := h.addressRepo.Create(address)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	claims, err := token.GetClaimsFromContext(r)
	if err != nil {
		http.Error(w, "invalid credentials", http.StatusUnauthorized)
		return
	}

	addresses, err := h.addressRepo.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for _, a := range addresses {
		saveID := a.ID
		a.ID = 0

		if *a == *address {
			a.ID = saveID
			address.ID = a.ID

			userAddress := &models.UsersAddressesIDPair{
				UserID:    claims.ID,
				AddressID: address.ID,
			}

			err = h.usersAddressesRepo.Create(userAddress)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			w.WriteHeader(http.StatusOK)
		}
	}

	http.Error(w, "", http.StatusBadRequest)
}

func (h *AddressHandler) GetUserAddresses(w http.ResponseWriter, r *http.Request) {
	claims, err := token.GetClaimsFromContext(r)
	if err != nil {
		http.Error(w, "invalid credentials", http.StatusUnauthorized)
		return
	}

	addresses, err := h.usersAddressesRepo.GetAddressesForUserByID(claims.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	addressesResponse := &responses.AddressCollectionResponse{}

	for _, a := range addresses {
		address := &responses.AddressResponse{
			ID:          a.ID,
			City:        a.City,
			Street:      a.Street,
			HouseNumber: a.HouseNumber,
			ZipCode:     a.ZipCode,
			FloorNumber: a.FloorNumber.String,
			Apartment:   a.Apartment.String,
		}
		addressesResponse.Addresses = append(addressesResponse.Addresses, address)
	}

	json.NewEncoder(w).Encode(addressesResponse)
	w.WriteHeader(http.StatusOK)
}

func (h *AddressHandler) UpdateAddress(w http.ResponseWriter, r *http.Request) {
	req := new(request.AddressRequest)
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	addresses, err := h.addressRepo.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	address := &models.Address{
		City:        req.City,
		Street:      req.Street,
		HouseNumber: req.HouseNumber,
		ZipCode:     req.ZipCode,
		FloorNumber: util.NullString(req.FloorNumber),
		Apartment:   util.NullString(req.Apartment),
	}

	for _, a := range addresses {
		savedID := a.ID
		a.ID = 0

		if *a == *address {
			a.ID = savedID

			err := h.addressRepo.Update(address)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			w.WriteHeader(http.StatusOK)
			return
		}
	}

	http.Error(w, "couldn't find address", http.StatusBadRequest)
}

func (h *AddressHandler) RemoveAddress(w http.ResponseWriter, r *http.Request) {
	req := new(request.AddressRequest)
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	address := &models.Address{
		City:        req.City,
		Street:      req.Street,
		HouseNumber: req.HouseNumber,
		ZipCode:     req.ZipCode,
		FloorNumber: util.NullString(req.FloorNumber),
		Apartment:   util.NullString(req.Apartment),
	}

	addresses, err := h.addressRepo.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	for _, a := range addresses {
		savedID := a.ID
		a.ID = 0

		if *a == *address {
			a.ID = savedID

			err := h.addressRepo.Delete(a.ID)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			w.WriteHeader(http.StatusOK)
			return
		}
	}
}

func (h *AddressHandler) RemoveUserAddress(w http.ResponseWriter, r *http.Request) {
	req := new(request.AddressRequest)
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	http.Error(w, "couldn't find address", http.StatusBadRequest)
}
