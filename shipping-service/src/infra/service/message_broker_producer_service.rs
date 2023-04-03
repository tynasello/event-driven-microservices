use crate::{
    application::interfaces::i_message_broker_producer_service::IMessageBrokerProducerService,
    infra::message_broker::kafka_config::KafkaProducer,
};

pub struct MessageBrokerProducerService<'a> {
    pub kafka_producer: &'a mut KafkaProducer,
}

impl<'a> IMessageBrokerProducerService for MessageBrokerProducerService<'a> {
    fn publish_message(&mut self, message: &str) {
        self.kafka_producer.send_to_kafka(message.to_string());
    }
}
