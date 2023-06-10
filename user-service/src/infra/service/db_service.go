package service

import (
	"os"

	"example.com/user-service/src/infra/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DbService struct {
	Db *gorm.DB
}

func NewDbService() *DbService {
	dbs := &DbService{}
	db, err := gorm.Open(mysql.Open(os.Getenv("DB_URI")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	dbs.Db = db

	return dbs
}

func (dbs DbService) RunDbMigrations() {
	dbs.Db.AutoMigrate(&model.UserModel{})
}
