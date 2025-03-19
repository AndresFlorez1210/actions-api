package main

import (
	"actions-back/config"
	"actions-back/internal"
	"actions-back/routes"
	"log"
	"net/http"
)

func main() {
	container := internal.NewContainer()
	router := routes.ConfigureRoutes(container)

	port := config.GetEnv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
