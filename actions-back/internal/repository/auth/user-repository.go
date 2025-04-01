package auth

import (
	"context"

	entity "actions-back/internal/entity/auth"

	"github.com/uptrace/bun"
)

type UserRepositoryImpl struct {
	db *bun.DB
}

func NewUserRepository(db *bun.DB) UserRepository {
	return &UserRepositoryImpl{db: db}
}

func (repository *UserRepositoryImpl) Create(ctx context.Context, user *entity.User) error {
	_, err := repository.db.NewInsert().Model(user).Exec(ctx)
	return err
}

func (repository *UserRepositoryImpl) GetByUsername(ctx context.Context, username string) (*entity.User, error) {
	user := new(entity.User)
	err := repository.db.NewSelect().Model(user).Where("username = ?", username).Scan(ctx)
	return user, err
}
