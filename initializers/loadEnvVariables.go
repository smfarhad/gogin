package initializers

import (
	"exapmle/gogin/initializers"
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	initializers.LoadEnvVariables()
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
