package usecase

import (
	"example.com/inventory-service/src/application/interfaces"
	"example.com/inventory-service/src/application/logic"
	"example.com/inventory-service/src/domain/entity"
)

type UpdateInventoryUseCase struct {
	InventoryRepository interfaces.IInventoryRepository
}

func (u UpdateInventoryUseCase) Execute(Inventory entity.Inventory, AddToQuantity int) *logic.Result[entity.Inventory] {
	// see if inventory exists
	inventoryExistsResult := u.InventoryRepository.GetByLabel(Inventory.Label)

	if inventoryExistsResult.IsFailure {
		return logic.FailedResult[entity.Inventory]("Inventory not found")
	}

	existingInventory := inventoryExistsResult.GetValue()

	// reserved inventory

	existingInventory.QuantityInStock += AddToQuantity

	updatedInventoryResult := u.InventoryRepository.Update(existingInventory)
	if updatedInventoryResult.IsFailure {
		return logic.FailedResult[entity.Inventory]("Failed to update inventory")
	}

	return updatedInventoryResult
}
