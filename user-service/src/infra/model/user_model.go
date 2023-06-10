package model

import "gorm.io/gorm"

type UserModel struct {
	gorm.Model
	Id       int    `gorm:"column:id;primaryKey"`
	Username string `gorm:"column:username;unique"`
	Password string `gorm:"column:password"`
}

func (UserModel) TableName() string {
	return "user"
}
