package service

import (
	"encoding/json"
	"fmt"

	"example.com/inventory-service/src/application/usecase"
	"example.com/inventory-service/src/domain/enum"
	"example.com/inventory-service/src/domain/event"
	"example.com/inventory-service/src/infra/messagebroker"
)

type MessageBrokerConsumerService struct {
	_orderEvent             event.OrderEvent
	ReserveInventoryUseCase usecase.ReserveInventoryUseCase
	OrderCancelledUseCase   usecase.OrderCancelledUseCase
	KafkaBroker             messagebroker.KafkaBroker
}

func (m MessageBrokerConsumerService) StartConsuming() {
	m.KafkaBroker.ConsumeFromKafka(m)
}

func (m MessageBrokerConsumerService) ConsumeMessage(message string) {
	err := json.Unmarshal([]byte(message), &m._orderEvent)
	if err != nil {
		fmt.Println(err)
		return
	}

	if !m._orderEvent.IsValidToConsume() {
		return
	}

	switch m._orderEvent.EventType {
	case enum.ORDER_REQUESTED.String():
		m.orderRequested()
	case enum.ORDER_CANCELLED.String():
		m.orderCancelled()
	}
	// else the order event type is not of concern to this service
}

func (m MessageBrokerConsumerService) orderRequested() {
	if m._orderEvent.OrderId == 0 || m._orderEvent.InventoryLabel == "" || m._orderEvent.InventoryQauntity == 0 {
		fmt.Println("Invalid ORDER_REQUESTED event")
		return
	}
	m.ReserveInventoryUseCase.Execute(m._orderEvent.OrderId, m._orderEvent.InventoryLabel, m._orderEvent.InventoryQauntity)
}

func (m MessageBrokerConsumerService) orderCancelled() {
	if m._orderEvent.OrderId == 0 || (m._orderEvent.IsInventoryReserved == true && m._orderEvent.InventoryQauntity == 0) || m._orderEvent.InventoryLabel == "" {
		fmt.Println("Invalid ORDER_CANCELLED event")
		return
	}
	m.OrderCancelledUseCase.Execute(m._orderEvent.IsInventoryReserved, m._orderEvent.InventoryLabel, m._orderEvent.InventoryQauntity)
}
