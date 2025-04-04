package routes

import (
	"actions-back/internal"

	"github.com/gorilla/mux"
)

func ConfigureRoutes(container *internal.Container) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/actions", container.ActionsController.GetActions).Methods("GET")
	router.HandleFunc("/actions/best-actions", container.ActionsController.GetBestActions).Methods("GET")

	router.HandleFunc("/auth/register", container.AuthController.Register).Methods("POST")
	router.HandleFunc("/auth/login", container.AuthController.Login).Methods("POST")

	router.HandleFunc("/actions/filter", container.ActionsController.FilterActionsByKeyword).Methods("GET")

	return router
}
