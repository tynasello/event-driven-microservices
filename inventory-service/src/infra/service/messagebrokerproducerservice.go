package service

import (
	"example.com/inventory-service/src/infra/messagebroker"
)

type MessageBrokerProducerService struct {
	KafkaBroker messagebroker.KafkaBroker
}

func (m MessageBrokerProducerService) PublishMessage(message string) {
	m.KafkaBroker.SendToKafka(message)
}
