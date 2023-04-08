from enum import Enum


class EOrderEventType(Enum):
    INVENTORY_RESERVED = 1
    TRANSACTION_COMPLETED = 2
    TRANSACTION_FAILED = 3
