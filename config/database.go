package config

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/Nitesh-04/locked-in/models"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	err = db.AutoMigrate(&models.User{}, &models.Group{}, &models.GroupMembership{}, &models.StudySession{})
	if err != nil {
		log.Fatal("Migration failed:", err)
	}

	DB = db
	fmt.Println("Database connected & migrated!")
}
