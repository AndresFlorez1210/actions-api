package auth

import (
	"context"

	entity "actions-back/internal/entity/auth"
)

type UserRepository interface {
	Create(ctx context.Context, user *entity.User) error
	GetByEmail(ctx context.Context, email string) (*entity.User, error)
}
