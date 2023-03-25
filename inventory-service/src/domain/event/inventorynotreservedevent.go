package event

import "example.com/inventory-service/src/domain/enum"

type InventoryNotReservedEvent struct {
	EventType string `json:"eventType"`
	OrderId   int    `json:"orderId"`
}

func NewInventoryNotReservedEvent(orderId int) InventoryNotReservedEvent {
	return InventoryNotReservedEvent{
		EventType: enum.INVENTORY_NOT_RESERVED.String(),
		OrderId:   orderId,
	}
}
