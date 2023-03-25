package interfaces

import (
	"example.com/inventory-service/src/application/logic"
	"example.com/inventory-service/src/domain/entity"
)

type IInventoryRepository interface {
	Create(inventory entity.Inventory) *logic.Result[entity.Inventory]
	GetById(id int) *logic.Result[entity.Inventory]
	GetByLabel(label string) *logic.Result[entity.Inventory]
	Update(inventory entity.Inventory) *logic.Result[entity.Inventory]
}
