package config

import (
	"log"
	"os"

	"example.com/inventory-service/src/internal/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func LoadEnvVariables() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

var Db *gorm.DB

func ConnectToDb() {
	dsn := os.Getenv("GORM_DB_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	Db = db
}

func RunDbMigrations() {
	Db.AutoMigrate(&models.InventoryItem{})
}
