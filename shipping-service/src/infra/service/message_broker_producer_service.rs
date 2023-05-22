use crate::{
    application::interfaces::i_message_broker_producer_service::IMessageBrokerProducerService,
    infra::message_broker::kafka_config::KafkaProducer,
};

pub struct MessageBrokerProducerService<'a> {
    kafka_producer: &'a mut KafkaProducer,
}

impl<'a> MessageBrokerProducerService<'a> {
    pub fn new(kafka_producer: &'a mut KafkaProducer) -> Self {
        Self { kafka_producer }
    }
}

impl IMessageBrokerProducerService for MessageBrokerProducerService<'_> {
    fn publish_message(&mut self, message: &str) {
        self.kafka_producer.send_to_kafka(message);
    }
}
