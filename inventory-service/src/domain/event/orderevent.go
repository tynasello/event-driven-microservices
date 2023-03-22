package event

import "example.com/inventory-service/src/domain/enum"

type OrderEvent struct {
	EventType       string
	OrderId         int
	ProductName     string
	ProductQuantity int
}

func (o OrderEvent) IsValidToConsume() bool {
	return o.EventType == enum.ORDER_REQUESTED.String() || o.EventType == enum.ORDER_CANCELLED.String()
}
