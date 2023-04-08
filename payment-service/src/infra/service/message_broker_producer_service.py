from application.interfaces.i_message_broker_producer_service import (
    IMessageBrokerProducerService,
)
from infra.message_broker.kafka_config import KafkaProducer


class MessageBrokerProducerService(IMessageBrokerProducerService):
    def __init__(self, kafka_producer: KafkaProducer):
        self.kafka_producer = kafka_producer

    def publish_message(self, message):
        self.kafka_producer.send_to_kafka(message)
