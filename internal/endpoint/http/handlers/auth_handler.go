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

	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	userRepo *db.UserRepository
	cfg      *config.Config
}

func NewAuthHandler(ur *db.UserRepository, cfg *config.Config) *AuthHandler {
	return &AuthHandler{
		userRepo: ur,
		cfg:      cfg,
	}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	req := new(request.LoginRequest)
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := h.userRepo.GetUserByEmail(req.Email)
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	tokenService := token.NewTokenService(h.cfg)

	accessString, err := tokenService.GenerateAccessToken(user.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	refreshString, err := tokenService.GenerateRefreshToken(user.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := responses.TokenPairResponse{
		AccessToken:  accessString,
		RefreshToken: refreshString,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func (h *AuthHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	req := new(request.SignUpRequest)
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if req.Password != req.PasswordAgain {
		http.Error(w, "Passwords doesn't match!", http.StatusBadRequest)
		return
	}

	users, err := h.userRepo.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for _, user := range users {
		if user.Email == req.Email {
			http.Error(w, "Email already in use!", http.StatusBadRequest)
			return
		}
	}

	password, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Couldn't generate password!", http.StatusBadRequest)
		return
	}

	user := &models.User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Password:  string(password),
	}

	h.userRepo.Create(user)

	w.WriteHeader(http.StatusOK)
}

func (h *AuthHandler) Refresh(w http.ResponseWriter, r *http.Request) {
	tokenService := token.NewTokenService(h.cfg)

	claims, err := tokenService.GetClaims(r, h.cfg.RefreshSecret)
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	accessString, err := tokenService.GenerateAccessToken(claims.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	refreshString, err := tokenService.GenerateRefreshToken(claims.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := responses.TokenPairResponse{
		AccessToken:  accessString,
		RefreshToken: refreshString,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func (h *AuthHandler) ResetPassword(w http.ResponseWriter, r *http.Request) {
	req := new(request.ResetPasswordRequest)
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
