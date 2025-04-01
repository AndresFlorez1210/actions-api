package authController

import (
	"actions-back/internal/services/auth"
	"encoding/json"
	"net/http"

	entity "actions-back/internal/entity/auth"
)

type AuthController struct {
	authService auth.AuthService
}

func NewAuthController(authService auth.AuthService) *AuthController {
	return &AuthController{authService: authService}
}

func (controller *AuthController) Register(response http.ResponseWriter, request *http.Request) {
	var user entity.User
	if err := json.NewDecoder(request.Body).Decode(&user); err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}
	token, err := controller.authService.Register(request.Context(), &user)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(response).Encode(map[string]string{"token": token})
}

func (controller *AuthController) Login(response http.ResponseWriter, request *http.Request) {
	var credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(request.Body).Decode(&credentials); err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}

	token, err := controller.authService.Login(request.Context(), credentials.Username, credentials.Password)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(response).Encode(map[string]string{"token": token})
}
