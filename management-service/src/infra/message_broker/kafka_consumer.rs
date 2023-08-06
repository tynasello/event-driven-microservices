use std::env;

use kafka::consumer::{Consumer, FetchOffset};

use crate::application::interfaces::i_message_broker_consumer::IMessageBrokerConsumer;

pub struct KafkaConsumer<'a> {
    kafka_consumer: &'a mut Consumer,
}

impl<'a> KafkaConsumer<'a> {
    pub fn new(kafka_consumer: &'a mut Consumer) -> Self {
        Self { kafka_consumer }
    }
}

impl<'a> IMessageBrokerConsumer for KafkaConsumer<'a> {
    fn consume_messages(&mut self) -> Result<Vec<String>, String> {
        let mut messages = vec![];
        messages.clear();
        for message_set in self.kafka_consumer.poll().unwrap().iter() {
            for raw_message in message_set.messages() {
                let message = std::str::from_utf8(raw_message.value).unwrap();
                messages.push(message.to_string());
            }
            self.kafka_consumer.consume_messageset(message_set).unwrap();
        }
        self.kafka_consumer.commit_consumed().unwrap();
        return Ok(messages);
    }
}

pub fn setup_kafka_consumer() -> Consumer {
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
