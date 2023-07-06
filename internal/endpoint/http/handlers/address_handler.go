package handler

import (
	config "hangryAPI/configs"
	"hangryAPI/internal/repositories/db"
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
