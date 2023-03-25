package usecase

import (
	"encoding/json"

	"example.com/inventory-service/src/application/interfaces"
	"example.com/inventory-service/src/application/logic"
	"example.com/inventory-service/src/domain/event"
)

type ReserveInventoryUseCase struct {
	InventoryRepository          interfaces.IInventoryRepository
	MessageBrokerProducerService interfaces.IMessageBrokerProducerService
}

func (u ReserveInventoryUseCase) Execute(orderId int, inventoryLabel string, inventoryQauntity int) *logic.Result[bool] {
	// see if inventory exists and can be reserved
	inventoryExistsResult := u.InventoryRepository.GetByLabel(inventoryLabel)
	if inventoryExistsResult.IsFailure || inventoryExistsResult.GetValue().QuantityReserved+inventoryQauntity > inventoryExistsResult.GetValue().QuantityInStock {
		inventoryNotReservedEvent := event.NewInventoryNotReservedEvent(orderId)
		inventoryNotReservedEventJson, err := json.Marshal(inventoryNotReservedEvent)
		if err != nil {
			return logic.FailedResult[bool](err.Error())
		}
		u.MessageBrokerProducerService.PublishMessage(string(inventoryNotReservedEventJson))
		return logic.FailedResult[bool]("Inventory not found")
	}

	existingInventory := inventoryExistsResult.GetValue()

	// reserved inventory

	existingInventory.QuantityReserved += inventoryQauntity

	updatedInventoryResult := u.InventoryRepository.Update(existingInventory)
	if updatedInventoryResult.IsFailure {
		return logic.FailedResult[bool]("Failed to update inventory")
	}

	inventoryReservedEvent := event.NewInventoryReservedEvent(orderId)
	inventoryReservedEventJson, err := json.Marshal(inventoryReservedEvent)
	if err != nil {
		return logic.FailedResult[bool](err.Error())
	}
	u.MessageBrokerProducerService.PublishMessage(string(inventoryReservedEventJson))

	return logic.OkResult(true)
}
