package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/nuck7/nfl-picks-server/src/api"
	"github.com/nuck7/nfl-picks-server/src/database"
	"github.com/nuck7/nfl-picks-server/src/models"
	"github.com/nuck7/nfl-picks-server/src/seed"
)

func main() {
	dot_env_err := godotenv.Load("config/.env.local")
	if dot_env_err != nil {
		log.Fatalf("Some error occured. Err: %s", dot_env_err)
	}

	config :=
		database.Config{
			Host:     os.Getenv("DB_HOST"),
			User:     os.Getenv("DB_USERNAME"),
			Password: os.Getenv("DB_PASSWORD"),
			DB:       os.Getenv("DB_NAME"),
		}
	connectionString := database.GetConnectionString(config)
	err := database.Connect(connectionString)
	if err != nil {
		panic(err.Error())
	}

	db := database.Connector
	// Run database migrations
	db.AutoMigrate(&models.Team{}, &models.User{}, &models.Pick{}, &models.Week{}, &models.Matchup{})

	// Seed database
	seed.LoadAll(db)

	// Start API
	api.Api()

}
