package initializers

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	// var err error
	// dsn := "host=localhost user=user password=fintech..2022 dbname=database port=5432 sslmode=disable"
	// DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	// if err != nil {
	// 	log.Fatal("Failed to connect to database")
	// }

	var err error

	// Fetching credentials from environment variables for security
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	dsn := "host=" + host + " user=" + user + " password=" + password + " dbname=" + dbname + " port=" + port + " sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("Database connection established")
}
