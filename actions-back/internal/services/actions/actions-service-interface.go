package services

import (
	entity "actions-back/internal/entity/actions"
	"context"
)

type ActionsService interface {
	GetActions(ctx context.Context) ([]entity.Action, error)
	GetBestActions(ctx context.Context) ([]entity.Action, error)
	FilterActionsByKeyword(ctx context.Context, requestFilter entity.FilterAction) ([]entity.Action, error)
}
