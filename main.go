package main

import (
	"log"
	"os"

	"github.com/UnwishingMoon/pixiu-bot/pkg/bot"
	"github.com/UnwishingMoon/pixiu-bot/pkg/db"
	"github.com/joho/godotenv"
)

func main() {

	// If not production read env variables from .env file
	if os.Getenv("GO_ENV") != "prod" {
		err := godotenv.Load()
		if err != nil {
			log.Fatalf("Error loading .env file: %s", err)
		}
	}

	// Starting the database connection
	db.Start()

	// Starting bot
	bot.Start()
}
