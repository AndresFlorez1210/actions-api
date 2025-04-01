package utils

import entity "actions-back/internal/entity/actions"

func FilterActions(actions []entity.Action, f func(action entity.Action) bool) []entity.Action {
	filteredActions := []entity.Action{}
	for _, action := range actions {
		if f(action) {
			filteredActions = append(filteredActions, action)
		}
	}
	return filteredActions
}
