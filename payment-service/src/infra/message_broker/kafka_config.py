from confluent_kafka import Consumer, Producer


class KafkaProducer:
    def __init__(self):
        self.producer = config_kafka_producer()

    def send_to_kafka(self, message):
        self.producer.produce("edms", message.encode("utf-8"))
        print("message sent to kafka: ", message)
        self.producer.flush()


class KafkaConsumer:
    def __init__(self):
        self.consumer = config_kafka_consumer()

    def consume_from_kafka(self):
        raw_message = self.consumer.poll(1.0)

        if raw_message is not None:
            message = raw_message.value().decode("utf-8")
            return message

        return

    def stop_consuming(self):
        self.consumer.close()


def config_kafka_producer():
    producer_config = {
        "bootstrap.servers": "kafka:9092",
    }
    return Producer(producer_config)


def config_kafka_consumer():
    consumer_config = {
        "bootstrap.servers": "kafka:9092",
        "group.id": "payment-service-consumer-group",
        "auto.offset.reset": "earliest",
    }

    consumer = Consumer(consumer_config)
    consumer.subscribe(["edms"])
    return consumer
