package event

import "example.com/inventory-service/src/domain/enum"

type OrderEvent struct {
	EventType           string `json:"eventType"`
	OrderId             int    `json:"orderId"`
	InventoryLabel      string `json:"inventoryLabel"`
	InventoryQuantity   int    `json:"inventoryQuantity"`
	IsInventoryReserved bool   `json:"isInventoryReserved"`
}

func (o OrderEvent) IsValidToConsume() bool {
	return o.EventType == enum.ORDER_REQUESTED.String() || o.EventType == enum.ORDER_CANCELLED.String()
}
