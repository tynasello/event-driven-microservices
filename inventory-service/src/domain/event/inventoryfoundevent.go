package event

import "example.com/inventory-service/src/domain/enum"

type InventoryFoundEvent struct {
	eventType string
	orderId   int
}

func NewInventoryFoundEvent(orderId int) InventoryFoundEvent {
	return InventoryFoundEvent{
		eventType: enum.INVENTORY_FOUND,
		orderId:   orderId,
	}
}
