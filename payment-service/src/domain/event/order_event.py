from domain.enums.e_order_event_type import EOrderEventType


class OrderEvent:
    def __init__(self, eventType: str, orderId: int):
        self.event_type = eventType
        self.order_id = orderId

    def isValidToConsume(self):
        return self.event_type == EOrderEventType.INVENTORY_RESERVED.name
