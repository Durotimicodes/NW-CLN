package database

import (
	"fmt"
	"log"
	"os"

	"github.com/durotimicodes/natwest-clone/user-service/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// InitDB initializes the PostGresSQL database connection
func InitDB() *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"))

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}

	//Auto-migrate the user model
	if err := db.AutoMigrate(&models.User{}); err != nil {
		log.Fatal("Failed to auto-migrate User model", err)
	}

	log.Println("Connecting to PostgreSQL DB successfully")

	return db
}
