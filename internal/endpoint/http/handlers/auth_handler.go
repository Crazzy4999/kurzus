package handler

import (
	config "hangryAPI/configs"
	repository "hangryAPI/internal/repositories"
	"net/http"
)

type AuthHandler struct {
	userRepo *repository.UserRepositoryI
	cfg      *config.Config
}

func NewAuthHandler(ur repository.UserRepositoryI, cfg *config.Config) *AuthHandler {
	return &AuthHandler{
		userRepo: &ur,
		cfg:      cfg,
	}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {

}

func (h *AuthHandler) SignUp(w http.ResponseWriter, r *http.Request) {

}

func (h *AuthHandler) Refresh(w http.ResponseWriter, r *http.Request) {

}
