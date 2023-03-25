package service

import (
	"os"

	"example.com/inventory-service/src/infra/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DbService struct {
	Db *gorm.DB
}

func NewDbService() *DbService {
	dbs := &DbService{}
	db, err := gorm.Open(postgres.Open(os.Getenv("DB_URI")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	dbs.Db = db

	return dbs
}

func (dbs DbService) RunDbMigrations() {
	dbs.Db.AutoMigrate(&model.InventoryModel{})
}
