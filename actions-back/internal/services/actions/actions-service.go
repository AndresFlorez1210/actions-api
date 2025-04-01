package services

import (
	entity "actions-back/internal/entity/actions"
	repository "actions-back/internal/repository/actions"
	"actions-back/internal/utils"
	"context"
)

const (
	BUY_RATING  = "Buy"
	SELL_RATING = "Sell"
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

func (service *ActionsServiceImpl) GetBestActions(ctx context.Context) ([]entity.Action, error) {
	actions, err := service.repository.GetActions(ctx)
	if err != nil {
		return nil, err
	}

	return GetBestActionsForCategory(actions)
}

func (service *ActionsServiceImpl) FilterActionsByKeyword(ctx context.Context, requestFilter entity.FilterAction) ([]entity.Action, error) {
	actions, err := service.repository.FilterActionsByKeyword(ctx, requestFilter)
	if err != nil {
		return nil, err
	}

	return actions, nil
}

func GetBestActionsForCategory(actions []entity.Action) ([]entity.Action, error) {
	actions = utils.FilterActions(actions, func(action entity.Action) bool {
		return action.RatingFrom == BUY_RATING && action.RatingTo == BUY_RATING
	})

	actions = utils.FilterActions(actions, func(action entity.Action) bool {
		targetTo := utils.StringMoneyToFloat(action.TargetTo)
		targetFrom := utils.StringMoneyToFloat(action.TargetFrom)

		return targetTo > targetFrom
	})

	if len(actions) > 3 {
		return actions[:3], nil
	}

	return actions, nil
}
