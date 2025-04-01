package controller

import (
	entity "actions-back/internal/entity/actions"
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

func (controller *ActionsController) GetBestActions(responseWriter http.ResponseWriter, request *http.Request) {
	actions, err := controller.service.GetBestActions(request.Context())
	if err != nil {
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(responseWriter).Encode(actions)
}

func (controller *ActionsController) FilterActionsByKeyword(responseWriter http.ResponseWriter, request *http.Request) {
	key := request.URL.Query().Get("key")
	value := request.URL.Query().Get("value")

	if key == "" || value == "" {
		http.Error(responseWriter, "key and value query parameters are required", http.StatusBadRequest)
		return
	}

	requestFilter := entity.FilterAction{
		Key:   key,
		Value: value,
	}

	actions, err := controller.service.FilterActionsByKeyword(request.Context(), requestFilter)
	if err != nil {
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(responseWriter).Encode(actions)
}
