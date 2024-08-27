package initializers

import (
	"Backend_Todo/models"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("todo.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v\n", err)
	}

	log.Println("Database connection established")

	// Migrate the schema
	DB.AutoMigrate(&models.Todo{})
}
