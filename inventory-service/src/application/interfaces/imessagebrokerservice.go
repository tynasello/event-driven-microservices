package interfaces

type IMessageBrokerService interface {
	PublishMessage(message string)
	ConsumeMessage(message string)
}
