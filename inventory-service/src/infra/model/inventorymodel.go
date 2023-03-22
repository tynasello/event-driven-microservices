package model

import "gorm.io/gorm"

type InventoryModel struct {
	gorm.Model
	Id               int
	Label            string
	QuantityInStock  int
	QuantityReserved int
}

func (InventoryModel) TableName() string {
	return "inventory"
}
