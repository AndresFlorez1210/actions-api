package auth

import (
	"actions-back/config"
	entity "actions-back/internal/entity/auth"
	authRepository "actions-back/internal/repository/auth"
	"context"
	"errors"
)

type AuthServiceImpl struct {
	userRepository authRepository.UserRepository
}

func NewAuthService(userRepository authRepository.UserRepository) AuthService {
	return &AuthServiceImpl{userRepository: userRepository}
}

func (service *AuthServiceImpl) Register(ctx context.Context, user *entity.User) (string, error) {
	if err := user.HashPassword(); err != nil {
		return "", err
	}
	if err := service.userRepository.Create(ctx, user); err != nil {
		return "", err
	}
	return config.GenerateJWT(user.ID)
}

func (service *AuthServiceImpl) Login(ctx context.Context, email, password string) (string, error) {
	user, err := service.userRepository.GetByEmail(ctx, email)
	if err != nil || !user.ComparePassword(password) {
		return "", errors.New("invalid credentials")
	}

	return config.GenerateJWT(user.ID)
}
