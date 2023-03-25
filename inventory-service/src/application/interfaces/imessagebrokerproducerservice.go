package interfaces

type IMessageBrokerProducerService interface {
	PublishMessage(message string)
}
