package model

import "gorm.io/gorm"

type UserModel struct {
	gorm.Model
	Id       int    `gorm:"primaryKey"`
	Username string `gorm:"unique"`
	Password string
}

func (UserModel) TableName() string {
	return "user"
}
