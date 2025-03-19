package services

import (
	entity "actions-back/internal/entity/actions"
	repository "actions-back/internal/repository/actions"
	"context"
)

type ActionsServiceImpl struct {
	repository repository.ActionsRepository
}

func NewActionsService(repository repository.ActionsRepository) ActionsService {
	return &ActionsServiceImpl{repository: repository}
}

func (service *ActionsServiceImpl) GetActions(ctx context.Context) ([]entity.Action, error) {
	return service.repository.GetActions(ctx)
}
