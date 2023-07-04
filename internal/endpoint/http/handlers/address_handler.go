package handler

import (
	"encoding/json"
	config "hangryAPI/configs"
	request "hangryAPI/internal/endpoint/http/requests"
	"hangryAPI/internal/models"
	"hangryAPI/internal/repositories/db"
	"hangryAPI/internal/util"
	"net/http"
)

type AddressHandler struct {
	addressRepo *db.AddressRepository
	cfg         *config.Config
}

func NewAddressHandler(addressRepo *db.AddressRepository, cfg *config.Config) *AddressHandler {
	return &AddressHandler{
		addressRepo: addressRepo,
		cfg:         cfg,
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

/*func (h *AddressHandler) GetAddresses(w http.ResponseWriter, r *http.Request) {
	claims, err := token.GetClaimsFromContext(r)
	if err != nil {
		http.Error(w, "invalid credentials", http.StatusUnauthorized)
		return
	}



	w.WriteHeader(http.StatusOK)
}*/

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

/*func (h *AddressHandler) RemoveAddress(w http.ResponseWriter, r *http.Request) {
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
}*/
