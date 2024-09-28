package initializers

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

const (
	host     = "localhost"
	user     = "postgres"
	password = "tremothegoat"
	dbName   = "delivery"
	port     = 5432
)

var DB *gorm.DB

func ConnectDb() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		host, user, password, dbName, port)

	// Assign the result to the global DB variable
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	// Handle the error properly
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("Connected to the database successfully!")
}
