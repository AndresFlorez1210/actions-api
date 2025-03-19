package controller

import (
	services "actions-back/internal/services/actions"
	"encoding/json"
	"net/http"
)

type ActionsController struct {
	service services.ActionsService
}

func NewActionsController(service services.ActionsService) *ActionsController {
	return &ActionsController{service: service}
}

func (controller *ActionsController) GetActions(responseWriter http.ResponseWriter, request *http.Request) {
	actions, err := controller.service.GetActions(request.Context())
	if err != nil {
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(responseWriter).Encode(actions)
}
