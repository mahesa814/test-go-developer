package main

import (
	"log"
	"test-go-developer/app/route"
	"test-go-developer/configs"
	"test-go-developer/database"
)

func main() {
	configs.LoadEnv()
	database.Connect()

	// Create extension default uuid
	if err := database.DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";").Error; err != nil {
		log.Fatalf("Failed to create extension \"uuid-ossp\": %v", err)
	}

	r := route.SetupRouter()
	if err := r.Run(":" + configs.AppPort); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
