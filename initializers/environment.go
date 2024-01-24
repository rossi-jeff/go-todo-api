package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvironment() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}
