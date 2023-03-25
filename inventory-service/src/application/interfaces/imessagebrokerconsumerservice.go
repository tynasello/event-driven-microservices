package interfaces

type IMessageBrokerConsumerService interface {
	ConsumeMessage(message string)
}
