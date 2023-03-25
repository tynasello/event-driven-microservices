package repository

import (
	"example.com/inventory-service/src/application/logic"
	"example.com/inventory-service/src/domain/entity"
	"example.com/inventory-service/src/infra/model"
	"gorm.io/gorm"
)

type InventoryRepository struct {
	Db *gorm.DB
}

func (r InventoryRepository) Create(inventory entity.Inventory) *logic.Result[entity.Inventory] {
	inventoryModel := model.InventoryModel{
		Label:            inventory.Label,
		QuantityInStock:  inventory.QuantityInStock,
		QuantityReserved: inventory.QuantityReserved,
	}
	inventoryCreatedResult := r.Db.Create(&inventoryModel)

	if inventoryCreatedResult.Error != nil {
		return logic.FailedResult[entity.Inventory]("Failed to create inventory item")
	}

	createdInventory := entity.Inventory{Id: inventoryModel.Id, Label: inventoryModel.Label, QuantityInStock: inventoryModel.QuantityInStock, QuantityReserved: inventoryModel.QuantityReserved}
	return logic.OkResult(createdInventory)
}

func (r InventoryRepository) GetById(id int) *logic.Result[entity.Inventory] {
	var inventoryModel model.InventoryModel

	existingInventoryResult := r.Db.First(&inventoryModel, id)

	if existingInventoryResult.Error != nil {
		return logic.FailedResult[entity.Inventory]("Failed to get inventory item")
	}

	existingInventory := entity.Inventory{Id: inventoryModel.Id, Label: inventoryModel.Label, QuantityInStock: inventoryModel.QuantityInStock, QuantityReserved: inventoryModel.QuantityReserved}
	return logic.OkResult(existingInventory)
}

func (r InventoryRepository) GetByLabel(label string) *logic.Result[entity.Inventory] {
	var inventoryModel model.InventoryModel

	existingInventoryResult := r.Db.Model(model.InventoryModel{Label: label}).First(&inventoryModel)

	if existingInventoryResult.Error != nil {
		return logic.FailedResult[entity.Inventory]("Failed to get inventory item")
	}

	existingInventory := entity.Inventory{Id: inventoryModel.Id, Label: inventoryModel.Label, QuantityInStock: inventoryModel.QuantityInStock, QuantityReserved: inventoryModel.QuantityReserved}
	return logic.OkResult(existingInventory)
}

func (r InventoryRepository) Update(inventory entity.Inventory) *logic.Result[entity.Inventory] {
	var inventoryModel model.InventoryModel

	existingInventoryResult := r.Db.First(&inventoryModel, inventory.Id)
	if existingInventoryResult.Error != nil {
		return logic.FailedResult[entity.Inventory]("Failed to get inventory item")
	}
	if inventoryModel.Id == 0 {
		return logic.FailedResult[entity.Inventory]("Inventory item not found")
	}

	updatedInventoryResult := r.Db.Model(&inventoryModel).Updates(
		model.InventoryModel{
			QuantityInStock:  inventory.QuantityInStock,
			QuantityReserved: inventory.QuantityReserved,
		})
	if updatedInventoryResult.Error != nil {
		return logic.FailedResult[entity.Inventory]("Failed to update inventory item")
	}

	createdInventory := entity.Inventory{Id: inventoryModel.Id, Label: inventoryModel.Label, QuantityInStock: inventoryModel.QuantityInStock, QuantityReserved: inventoryModel.QuantityReserved}
	return logic.OkResult(createdInventory)
}
