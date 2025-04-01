package main

import (
	"actions-back/config"
	"actions-back/internal"
	"actions-back/routes"
	"log"
	"net/http"

	"github.com/rs/cors"
)

func main() {

	container := internal.NewContainer()
	router := routes.ConfigureRoutes(container)

	corsConfig := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:8081"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
	})

	handlerRouter := corsConfig.Handler(router)

	port := config.GetEnv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	if err := http.ListenAndServe(":"+port, handlerRouter); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
