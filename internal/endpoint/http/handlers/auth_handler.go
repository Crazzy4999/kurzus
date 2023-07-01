package handler

import (
	"encoding/json"
	config "hangryAPI/configs"
	request "hangryAPI/internal/endpoint/http/requests"
	"hangryAPI/internal/endpoint/http/responses"
	"hangryAPI/internal/models"
	"hangryAPI/internal/repositories/db"
	"hangryAPI/internal/service/email"
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
		http.Error(w, "passwords doesn't match", http.StatusBadRequest)
		return
	}

	users, err := h.userRepo.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for _, user := range users {
		if user.Email == req.Email {
			http.Error(w, "email already in use", http.StatusBadRequest)
			return
		}
	}

	password, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "couldn't generate password", http.StatusBadRequest)
		return
	}

	user := &models.User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Password:  string(password),
	}

	if err = h.userRepo.Create(user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *AuthHandler) Refresh(w http.ResponseWriter, r *http.Request) {
	tokenService := token.NewTokenService(h.cfg)

	claims, err := token.GetClaimsFromContext(r)
	if err != nil {
		http.Error(w, "invalid credentials", http.StatusUnauthorized)
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

func (h *AuthHandler) GetPasswordResetKey(w http.ResponseWriter, r *http.Request) {
	req := new(request.GetResetKeyRequest)
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := h.userRepo.GetUserByEmail(req.Email)
	if err != nil {
		http.Error(w, "user doesn't exist with this email", http.StatusBadRequest)
		return
	}

	tokenService := token.NewTokenService(h.cfg)

	resetString, err := tokenService.GenerateRefreshToken(user.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	emailService := email.NewEmailService(h.cfg, "smtp.gmail.com", "587")
	err = emailService.SendResetEmail(req.Email, "Subject: Password reset for Hangry Food Delivery\n", "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n", "internal/service/email/reset_password_email.html", resetString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *AuthHandler) ResetPassword(w http.ResponseWriter, r *http.Request) {
	req := new(request.ResetPasswordRequest)
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if req.Password != req.PasswordAgain {
		http.Error(w, "passwords doesn't match", http.StatusBadRequest)
		return
	}

	user, err := h.userRepo.GetUserByEmail(req.Email)
	if err != nil {
		http.Error(w, "getting user failed", http.StatusBadRequest)
		return
	}

	resetKey := r.URL.Query().Get("reset_key")
	if err != nil {
		http.Error(w, "bad url, missing reset_key", http.StatusBadRequest)
		return
	}

	tokenService := token.NewTokenService(h.cfg)
	// TODO: fix invalid signature (invalid here in go, on jwt website it's fine for some reason)
	claims, err := tokenService.ValidateResetToken(resetKey)
	if err != nil {
		http.Error(w, "invalid credentials claims", http.StatusUnauthorized)
		return
	}

	if user.ID != claims.ID {
		http.Error(w, "invalid credentials", http.StatusUnauthorized)
		return
	}

	password, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "couldn't generate password", http.StatusBadRequest)
		return
	}

	user.Password = string(password)

	err = h.userRepo.Update(user)
	if err != nil {
		http.Error(w, "reseting password for user failed", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
