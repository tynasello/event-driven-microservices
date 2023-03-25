package model

import "gorm.io/gorm"

type InventoryModel struct {
	gorm.Model
	Id               int    `gorm:"primaryKey"`
	Label            string `gorm:"unique"`
	QuantityInStock  int
	QuantityReserved int
}

func (InventoryModel) TableName() string {
	return "inventory"
}
