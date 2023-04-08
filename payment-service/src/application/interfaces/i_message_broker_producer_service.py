from abc import ABC, abstractmethod


class IMessageBrokerProducerService(ABC):
    @abstractmethod
    def publish_message(self, message: str) -> None:
        pass
