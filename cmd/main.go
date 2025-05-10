package main

import (
	"auth-service/internal/config"
	"auth-service/internal/db"
	"auth-service/internal/routes"
	"auth-service/internal/services"
	"log"
	"net/http"
)

func main() {
	config.LoadConfig()
	db.ConnectDatabse()

	// add default user
	services.SeedDefault()
	router := routes.SetupRoutes()

	log.Printf("🚀 Server running on http://localhost:%s", config.AppConfig.Port)
	log.Fatal(http.ListenAndServe(":"+config.AppConfig.Port, router))
}
