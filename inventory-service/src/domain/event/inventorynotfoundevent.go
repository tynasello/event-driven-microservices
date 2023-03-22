package event

import "example.com/inventory-service/src/domain/enum"

type InventoryNotFoundEvent struct {
	eventType string
	orderId   int
}

func NewInventoryNotFoundEvent(orderId int) InventoryFoundEvent {
	return InventoryFoundEvent{
		eventType: enum.INVENTORY_NOT_FOUND,
		orderId:   orderId,
	}
}
