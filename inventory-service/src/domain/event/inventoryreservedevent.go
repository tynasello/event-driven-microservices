package event

import "example.com/inventory-service/src/domain/enum"

type InventoryReservedEvent struct {
	EventType string `json:"eventType"`
	OrderId   int    `json:"orderId"`
}

func NewInventoryReservedEvent(orderId int) InventoryReservedEvent {
	return InventoryReservedEvent{
		EventType: enum.INVENTORY_RESERVED.String(),
		OrderId:   orderId,
	}
}
