package main

import (
	"log"
	"test-go-developer/app/route"
	"test-go-developer/configs"
	"test-go-developer/database"
	"test-go-developer/database/migration"
	"test-go-developer/database/seeder"
)

func main() {
	configs.LoadEnv()
	database.Connect()
	if err := migration.AutoMigrate(); err != nil {
		log.Fatal(err)
	}

	seeder.SeedData(database.DB)

	r := route.SetupRouter()
	if err := r.Run(":" + configs.AppPort); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
