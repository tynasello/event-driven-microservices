use kafka::{
    consumer::{Consumer, FetchOffset},
    producer::{Producer, Record},
};
use std::env;

pub struct KafkaProducer {
    producer: Producer,
}

pub struct KafkaConsumer {
    consumer: Consumer,
}

impl KafkaProducer {
    pub fn new() -> Self {
        Self {
            producer: config_kafka_producer(),
        }
    }

    pub fn send_to_kafka(&mut self, message: &str) {
        let record = Record::from_value("edms", message.as_bytes());
        let message_sent_result = self.producer.send(&record);

        match message_sent_result {
            Ok(_) => println!("Message sent to Kafka"),
            Err(e) => println!("Error sending message to Kafka: {}", e),
        }
    }
}

impl KafkaConsumer {
    pub fn new() -> Self {
        Self {
            consumer: config_kafka_consumer(),
        }
    }

    pub fn consume_from_kafka(&mut self) -> Vec<String> {
        let mut messages = vec![];
        messages.clear();
        for message_set in self.consumer.poll().unwrap().iter() {
            for raw_message in message_set.messages() {
                let message = std::str::from_utf8(raw_message.value).unwrap();
                messages.push(message.to_string());
            }
            self.consumer.consume_messageset(message_set).unwrap();
        }
        self.consumer.commit_consumed().unwrap();
        return messages;
    }
}

fn config_kafka_producer() -> Producer {
    let hosts = get_host();
    let producer_result = Producer::from_hosts(hosts).create();
    match producer_result {
        Ok(producer) => producer,
        Err(e) => panic!("Error creating producer: {}", e),
    }
}

fn config_kafka_consumer() -> Consumer {
    let hosts = get_host();
    let topic = "edms";
    let consumer_result = Consumer::from_hosts(hosts)
        .with_topic(topic.to_string())
        .with_fallback_offset(FetchOffset::Latest)
        .with_group("shipping-service-consumer-group".to_owned())
        .create();
    match consumer_result {
        Ok(consumer) => consumer,
        Err(e) => panic!("Error creating consumer: {}", e),
    }
}

fn get_host() -> Vec<String> {
    let kafka_bootstrap_address = env::var("KAFKA_BOOTSTRAP_ADDRESS").unwrap();
    vec![kafka_bootstrap_address]
}
