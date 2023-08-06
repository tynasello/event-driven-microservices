from application.usecase.attempt_transaction_usecase import AttemptTransactionUseCase
from domain.enums.e_order_event_type import EOrderEventType
from domain.event.order_event import OrderEvent
from infra.message_broker.kafka_config import KafkaConsumer
import json


class MessageBrokerConsumerService:
    def __init__(
        self,
        kafka_consumer: KafkaConsumer,
        attempt_transaction_usecase: AttemptTransactionUseCase,
    ):
        self.kafka_consumer = kafka_consumer
        self.attempt_transaction_usecase = attempt_transaction_usecase

    def start_consuming(self):
        while True:
            message = self.kafka_consumer.consume_from_kafka()
            if message is not None:
                try:
                    self.orderEvent = OrderEvent(**json.loads(message))
                except Exception:
                    continue

                if not self.orderEvent.isValidToConsume():
                    continue

                match self.orderEvent.event_type:
                    case EOrderEventType.INVENTORY_RESERVED.name:
                        self.on_inventory_reserved()
                    case _:
                        continue

        self.kafka_consumer.stop_consuming()

    def on_inventory_reserved(self):
        if not type(self.orderEvent.order_id) is int:
            return
        self.attempt_transaction_usecase.execute(self.orderEvent.order_id)
