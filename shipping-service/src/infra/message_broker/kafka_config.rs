use kafka::{
    consumer::{Consumer, FetchOffset},
    producer::{Producer, Record},
};

pub struct KafkaProducer {
    producer: Producer,
}

pub struct KafkaConsumer {
    consumer: Consumer,
}

impl<'a> KafkaProducer {
    pub fn new() -> Self {
        Self {
            producer: config_kafka_producer(),
        }
    }

    pub fn send_to_kafka(&mut self, message: String) {
        let record = Record::from_value("edms", message.as_bytes());
        let message_sent_result = self.producer.send(&record);

        match message_sent_result {
            Ok(_) => println!("Message sent to Kafka"),
            Err(e) => println!("Error sending message to Kafka: {}", e),
        }
    }
}

impl<'a> KafkaConsumer {
    pub fn new() -> Self {
        Self {
            consumer: config_kafka_consumer(),
        }
    }

    pub fn consume_from_kafka(&mut self) -> Vec<String> {
        let mut messages = vec![];
        messages.clear();
        for ms in self.consumer.poll().unwrap().iter() {
            for raw_message in ms.messages() {
                let message = std::str::from_utf8(raw_message.value).unwrap();
                messages.push(message.to_string());
            }
            self.consumer.consume_messageset(ms).unwrap();
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
    let topic = "edms".to_string();
    let consumer_result = Consumer::from_hosts(hosts)
        .with_topic(topic)
        .with_fallback_offset(FetchOffset::Latest)
        .with_group("shipping-service-consumer-group".to_owned())
        .create();
    match consumer_result {
        Ok(consumer) => consumer,
        Err(e) => panic!("Error creating consumer: {}", e),
    }
}

fn get_host() -> Vec<String> {
    vec!["kafka:9092".to_string()]
}
