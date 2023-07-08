package handler

import (
	"encoding/json"
	config "hangryAPI/configs"
	request "hangryAPI/internal/endpoint/http/requests"
	"hangryAPI/internal/endpoint/http/responses"
	"hangryAPI/internal/models"
	"hangryAPI/internal/repositories/db"
	"hangryAPI/internal/service/token"
	"net/http"
)

type UserHandler struct {
	userRepo    *db.UserRepository
	addressRepo *db.AddressRepository
	cfg         *config.Config
}

func NewUserHandler(userRepo *db.UserRepository, addressRepo *db.AddressRepository, cfg *config.Config) *UserHandler {
	return &UserHandler{
		userRepo:    userRepo,
		addressRepo: addressRepo,
		cfg:         cfg,
	}
}

func (h *UserHandler) GetProfile(w http.ResponseWriter, r *http.Request) {
	tokenService := token.NewTokenService(h.cfg)
	claims, err := tokenService.GetClaims(r, h.cfg.AccessSecret)
	if err != nil {
		http.Error(w, "invalid credentials", http.StatusUnauthorized)
		return
	}

	user, err := h.userRepo.GetUserByID(claims.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp := responses.UserResponse{
		ID:    user.ID,
		Name:  user.FirstName + " " + user.LastName,
		Email: user.Email,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func (h *UserHandler) UpdateProfile(w http.ResponseWriter, r *http.Request) {
	req := new(request.UserRequest)
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	tokenService := token.NewTokenService(h.cfg)
	claims, err := tokenService.GetClaims(r, h.cfg.AccessSecret)
	if err != nil {
		http.Error(w, "invalid credentials", http.StatusUnauthorized)
		return
	}

	user, err := h.userRepo.GetUserByID(claims.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updatedUser := models.User{
		ID:        user.ID,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     user.Email,
		Password:  user.Password,
	}

	err = h.userRepo.Update(&updatedUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *UserHandler) DeleteProfile(w http.ResponseWriter, r *http.Request) {
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
		if a.UserID == claims.ID {
			err = h.addressRepo.Delete(a.ID)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
		}
	}

	err = h.userRepo.Delete(claims.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
