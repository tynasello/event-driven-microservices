import json
from domain.enums.e_order_event_type import EOrderEventType


class TransactionCompletedEvent:
    def __init__(self, order_id: int):
        self.event_type = EOrderEventType.TRANSACTION_COMPLETED.name
        self.order_id = order_id

    def to_json(self) -> str:
        return json.dumps({"eventType": self.event_type, "orderId": self.order_id})
