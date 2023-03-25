package usecase

import (
	"fmt"

	"example.com/inventory-service/src/application/interfaces"
	"example.com/inventory-service/src/application/logic"
)

type OrderCancelledUseCase struct {
	InventoryRepository          interfaces.IInventoryRepository
	MessageBrokerProducerService interfaces.IMessageBrokerProducerService
}

func (u OrderCancelledUseCase) Execute(isInventoryReserved bool, inventoryLabel string, inventoryQuantity int) *logic.Result[bool] {
	if !isInventoryReserved {
		return logic.OkResult(true)
	}
	inventoryExistsResult := u.InventoryRepository.GetByLabel(inventoryLabel)
	if inventoryExistsResult.IsFailure {
		fmt.Println(inventoryExistsResult.GetErrorMessage())
		return logic.FailedResult[bool]("Failed to get inventory")
	}

	existingInventory := inventoryExistsResult.GetValue()

	// free inventory
	existingInventory.QuantityReserved -= inventoryQuantity

	updatedInventoryResult := u.InventoryRepository.Update(existingInventory)
	if updatedInventoryResult.IsFailure {
		return logic.FailedResult[bool]("Failed to update inventory")
	}

	return logic.OkResult(true)
}
