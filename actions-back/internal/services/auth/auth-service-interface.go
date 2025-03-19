package auth

import (
	"context"

	entity "actions-back/internal/entity/auth"
)

type AuthService interface {
	Login(ctx context.Context, email, password string) (string, error)
	Register(ctx context.Context, user *entity.User) (string, error)
}
