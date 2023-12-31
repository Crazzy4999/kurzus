package handler

import (
	"encoding/json"
	config "hangryAPI/configs"
	request "hangryAPI/internal/endpoint/http/requests/auth"
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
		http.Error(w, "invalid credentials", http.StatusUnauthorized)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		http.Error(w, "invalid credentials", http.StatusUnauthorized)
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
		http.Error(w, JSON_TRANSFORM_FAILED, http.StatusBadRequest)
		return
	}

	if req.Password != req.PasswordAgain {
		http.Error(w, PASSWORD_MISMATCH, http.StatusBadRequest)
		return
	}

	users, err := h.userRepo.GetAll()
	if err != nil {
		http.Error(w, GET_ALL_USERS_FAILED, http.StatusBadRequest)
		return
	}

	for _, user := range users {
		if user.Email == req.Email {
			http.Error(w, EMAIL_TAKEN, http.StatusBadRequest)
			return
		}
	}

	password, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, PASSWORD_GENERATION_FAILED, http.StatusBadRequest)
		return
	}

	user := &models.User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Password:  string(password),
	}

	if err = h.userRepo.Create(user); err != nil {
		http.Error(w, CREATING_USER_FAILED, http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *AuthHandler) Refresh(w http.ResponseWriter, r *http.Request) {
	tokenService := token.NewTokenService(h.cfg)

	claims, err := tokenService.GetClaims(r, h.cfg.RefreshSecret)
	if err != nil {
		http.Error(w, INVALID_CREDENTIALS, http.StatusUnauthorized)
		return
	}

	accessString, err := tokenService.GenerateAccessToken(claims.ID)
	if err != nil {
		http.Error(w, TOKEN_GENERATION_FAILED, http.StatusInternalServerError)
		return
	}

	refreshString, err := tokenService.GenerateRefreshToken(claims.ID)
	if err != nil {
		http.Error(w, TOKEN_GENERATION_FAILED, http.StatusInternalServerError)
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
		http.Error(w, JSON_TRANSFORM_FAILED, http.StatusBadRequest)
		return
	}

	user, err := h.userRepo.GetUserByEmail(req.Email)
	if err != nil {
		http.Error(w, NO_SUCH_USER, http.StatusBadRequest)
		return
	}

	tokenService := token.NewTokenService(h.cfg)

	resetString, err := tokenService.GenerateResetToken(user.ID)
	if err != nil {
		http.Error(w, TOKEN_GENERATION_FAILED, http.StatusInternalServerError)
		return
	}

	emailService := email.NewEmailService(h.cfg, "smtp.gmail.com", "587")
	err = emailService.SendResetEmail(req.Email, "Subject: Password reset for Hangry Food Delivery\n", "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n", "internal/service/email/reset_password_email.html", resetString)
	if err != nil {
		http.Error(w, SEND_RESET_FAILED, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *AuthHandler) ResetPassword(w http.ResponseWriter, r *http.Request) {
	req := new(request.ResetPasswordRequest)
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, JSON_TRANSFORM_FAILED, http.StatusBadRequest)
		return
	}

	if req.Password != req.PasswordAgain {
		http.Error(w, PASSWORD_MISMATCH, http.StatusBadRequest)
		return
	}

	resetKey := r.URL.Query().Get("reset_key")
	if resetKey == "" {
		http.Error(w, MISSING_RESET_KEY, http.StatusBadRequest)
		return
	}

	tokenService := token.NewTokenService(h.cfg)
	claims, err := tokenService.ValidateResetToken(resetKey)
	if err != nil {
		http.Error(w, INVALID_CREDENTIALS, http.StatusUnauthorized)
		return
	}

	user, err := h.userRepo.GetUserByID(claims.ID)
	if err != nil {
		http.Error(w, GET_USER_FAILED, http.StatusBadRequest)
		return
	}

	if user.ID != claims.ID {
		http.Error(w, INVALID_CREDENTIALS, http.StatusUnauthorized)
		return
	}

	password, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, PASSWORD_GENERATION_FAILED, http.StatusBadRequest)
		return
	}

	user.Password = string(password)

	err = h.userRepo.Update(user)
	if err != nil {
		http.Error(w, PASSWORD_RESET_FAILED, http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
