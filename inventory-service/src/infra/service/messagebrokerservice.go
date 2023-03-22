package service

import (
	"encoding/json"
	"fmt"

	"example.com/inventory-service/src/application/usecase"
	"example.com/inventory-service/src/domain/enum"
	"example.com/inventory-service/src/domain/event"
	"example.com/inventory-service/src/infra/messagebroker"
)

type MessageBrokerService struct {
	_orderEvent              event.OrderEvent
	_reserveInventoryUseCase usecase.ReserveInventoryUseCase
	_orderCancelledUseCase   usecase.OrderCancelledUseCase
	KafkaBroker              messagebroker.KafkaBroker
}

func (m MessageBrokerService) PublishMessage(message string) {
	m.KafkaBroker.SendToKafka(message)
}

func (m MessageBrokerService) StartConsuming() {
	m.KafkaBroker.ConsumeFromKafka(m)
}

func (m MessageBrokerService) ConsumeMessage(message string) {
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

func (m MessageBrokerService) orderRequested() {
	reserveInventoryResult := m._reserveInventoryUseCase.Execute()
	if reserveInventoryResult.IsFailure {
		inventoryNotFoundEvent := event.NewInventoryNotFoundEvent(m._orderEvent.OrderId)
		inventoryNotFoundEventJson, err := json.Marshal(inventoryNotFoundEvent)
		if err != nil {
			fmt.Println(err)
			return
		}
		m.PublishMessage(string(inventoryNotFoundEventJson))
	} else {
		inventoryFoundEvent := event.NewInventoryFoundEvent(m._orderEvent.OrderId)
		inventoryFoundEventJson, err := json.Marshal(inventoryFoundEvent)
		if err != nil {
			fmt.Println(err)
			return
		}
		m.PublishMessage(string(inventoryFoundEventJson))
	}
}

func (m MessageBrokerService) orderCancelled() {
	orderCancelledResult := m._orderCancelledUseCase.Execute()
	if orderCancelledResult.IsFailure {
		fmt.Println(orderCancelledResult.GetErrorMessage())
	}
	fmt.Printf("Order %v cancelled\n", m._orderEvent.OrderId)
}
