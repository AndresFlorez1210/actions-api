package auth

import (
	"context"

	entity "actions-back/internal/entity/auth"
)

type UserRepository interface {
	Create(ctx context.Context, user *entity.User) error
	GetByUsername(ctx context.Context, username string) (*entity.User, error)
}
