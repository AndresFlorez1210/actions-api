package services

import (
	entity "actions-back/internal/entity/actions"
	"context"
)

type ActionsService interface {
	GetActions(ctx context.Context) ([]entity.Action, error)
}
