package repository

import (
	entity "actions-back/internal/entity/actions"
	"context"
)

type ActionsRepository interface {
	GetActions(ctx context.Context) ([]entity.Action, error)
	FilterActionsByKeyword(ctx context.Context, requestFilter entity.FilterAction) ([]entity.Action, error)
}
