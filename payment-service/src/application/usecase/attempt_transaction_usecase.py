import random
from application.interfaces.i_message_broker_producer_service import (
    IMessageBrokerProducerService,
)
from domain.event.transaction_completed_event import TransactionCompletedEvent
from domain.event.transaction_failed_event import TransactionFailedEvent


class AttemptTransactionUseCase:
    def __init__(self, message_broker_producer_service: IMessageBrokerProducerService):
        self.message_broker_producer_service = message_broker_producer_service

    def execute(self, order_id: int):
        ## add business specific logic here
        transaction_completed = random.random() < 0.9

        if transaction_completed:
            transaction_completed_event = TransactionCompletedEvent(order_id)
            self.message_broker_producer_service.publish_message(
                transaction_completed_event.to_json()
            )
        else:
            transaction_failed_event = TransactionFailedEvent(order_id)
            self.message_broker_producer_service.publish_message(
                transaction_failed_event.to_json()
            )
