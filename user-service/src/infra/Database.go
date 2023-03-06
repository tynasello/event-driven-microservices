package infra

import (
	"log"

	"example.com/user-service/src/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect(dbUri string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(dbUri), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to MySQL database")
	}
	log.Println("Connected to database")
	return db
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
	log.Println("Database migration completed")
}
