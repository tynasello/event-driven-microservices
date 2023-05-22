use kafka::consumer::{Consumer, FetchOffset};
use std::env;

pub struct KafkaConsumer {
    consumer: Consumer,
}

impl KafkaConsumer {
    pub fn new() -> Self {
        Self {
            consumer: config_kafka_consumer(),
        }
    }

    pub fn consume_from_kafka(&mut self) -> Vec<String> {
        let consumer = &mut self.consumer;
        let mut messages = vec![];
        messages.clear();
        for message_set in consumer.poll().unwrap().iter() {
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

fn config_kafka_consumer() -> Consumer {
    let hosts = get_host();
    let topic = "edms";
    let consumer_result = Consumer::from_hosts(hosts)
        .with_topic(topic.to_string())
        .with_fallback_offset(FetchOffset::Latest)
        .with_group("shipping-service-consumer-group".to_string())
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
