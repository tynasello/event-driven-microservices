from application.usecase.attempt_transaction_usecase import AttemptTransactionUseCase
from infra.message_broker.kafka_config import KafkaConsumer, KafkaProducer
from infra.service.message_broker_consumer_service import MessageBrokerConsumerService
from infra.service.message_broker_producer_service import MessageBrokerProducerService


def main():
    kafka_consumer = KafkaConsumer()
    kafka_producer = KafkaProducer()

    message_broker_producer_service = MessageBrokerProducerService(kafka_producer)

    attempt_transaction_usecase = AttemptTransactionUseCase(
        message_broker_producer_service
    )
    message_broker_consumer_service = MessageBrokerConsumerService(
        kafka_consumer, attempt_transaction_usecase
    )

    message_broker_consumer_service.start_consuming()


if __name__ == "__main__":
    main()
