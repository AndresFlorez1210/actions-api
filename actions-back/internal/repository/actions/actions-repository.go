package repository

import (
	entity "actions-back/internal/entity/actions"
	"context"

	"github.com/uptrace/bun"
)

type ActionsRepositoryImpl struct {
	db *bun.DB
}

func NewActionsRepository(db *bun.DB) ActionsRepository {
	return &ActionsRepositoryImpl{db: db}
}

func (repository *ActionsRepositoryImpl) GetActions(ctx context.Context) ([]entity.Action, error) {
	var actions []entity.Action
	err := repository.db.NewSelect().Model(&actions).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return actions, nil
}

func (repository *ActionsRepositoryImpl) FilterActionsByKeyword(ctx context.Context, requestFilter entity.FilterAction) ([]entity.Action, error) {
	var actions []entity.Action
	err := repository.db.NewSelect().Model(&actions).Where(requestFilter.Key+" = ?", requestFilter.Value).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return actions, nil
}
