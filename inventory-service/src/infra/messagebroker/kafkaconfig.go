package messagebroker

import (
	"fmt"
	"os"

	"example.com/inventory-service/src/application/interfaces"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaBroker struct {
	Producer *kafka.Producer
	Consumer *kafka.Consumer
}

func NewKafkaBroker() KafkaBroker {
	return KafkaBroker{
		Producer: ConfigKafkaProducer(),
		Consumer: ConfigKafkaConsumer(),
	}
}

func ConfigKafkaProducer() *kafka.Producer {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "kafka:9092",
		"client.id":         "edms-group-1",
		"acks":              "all",
	})
	if err != nil {
		fmt.Printf("Failed to create producer: %s\n", err)
	}
	return p
}

func ConfigKafkaConsumer() *kafka.Consumer {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "kafka:9092",
		"group.id":          "edms-group-1",
	})
	if err != nil {
		fmt.Printf("Failed to configure c: %s\n", err)
	}
	return c
}

func (k KafkaBroker) SendToKafka(message string) {
	topic := "edms"
	delivery_chan := make(chan kafka.Event, 10000)
	err := k.Producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(message),
	},
		delivery_chan,
	)
	if err != nil {
		fmt.Printf("Failed to produce message: %s\n", err)
	}
	fmt.Println("Message sent to kafka")
}

func (k KafkaBroker) ConsumeFromKafka(messageBrokerService interfaces.IMessageBrokerConsumerService) {
	topics := []string{"edms"}
	err := k.Consumer.SubscribeTopics(topics, nil)
	run := err == nil

	for run == true {
		ev := k.Consumer.Poll(10000)
		switch e := ev.(type) {
		case *kafka.Message:
			messageBrokerService.ConsumeMessage(string(e.Value))
		case kafka.Error:
			fmt.Fprintf(os.Stderr, "%% Error: %v\n", e)
			run = false
		}
	}

	k.Consumer.Close()
}
