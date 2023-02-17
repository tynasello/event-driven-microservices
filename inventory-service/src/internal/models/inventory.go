package models

import "gorm.io/gorm"

type InventoryItem struct {
	gorm.Model
	Label            string
	QuantityInStock  int
	QuantityReserved int
}

func (InventoryItem) TableName() string {
	return "inventory"
}
