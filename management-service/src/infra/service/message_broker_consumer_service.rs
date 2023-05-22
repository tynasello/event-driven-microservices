use crate::{
    application::interfaces::i_message_broker_consumer_service::IMessageBrokerConsumerService,
    infra::message_broker::kafka_config::KafkaConsumer,
};

pub struct MessageBrokerConsumerService<'a> {
    pub kafka_consumer: &'a mut KafkaConsumer,
}

impl IMessageBrokerConsumerService for MessageBrokerConsumerService<'_> {
    fn start_consuming(&mut self) {
        loop {
            let messages = self.kafka_consumer.consume_from_kafka();
            if messages.is_empty() {
                continue;
            }

            for message in messages {
                println!("{}", message);
            }
        }
    }
}
