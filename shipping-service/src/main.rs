use application::interfaces::i_message_broker_producer_service::IMessageBrokerProducerService;
use dotenv::dotenv;

use crate::application::interfaces::i_message_broker_consumer_service::IMessageBrokerConsumerService;
use crate::{
    application::usecase::order_accepted_usecase::OrderAcceptedUsecase,
    infra::{
        message_broker::kafka_config::{KafkaConsumer, KafkaProducer},
        service::{
            message_broker_consumer_service::MessageBrokerConsumerService,
            message_broker_producer_service::MessageBrokerProducerService,
        },
    },
};

mod application;
mod domain;
mod infra;

fn main() {
    dotenv().ok();

    let kafka_producer = &mut KafkaProducer::new();
    let kafka_consumer = &mut KafkaConsumer::new();

    let message_broker_producer_service: &mut dyn IMessageBrokerProducerService =
        &mut MessageBrokerProducerService::new(kafka_producer);

    let order_accepted_use_case = &mut OrderAcceptedUsecase::new(message_broker_producer_service);

    let mut message_broker_consumer_service =
        MessageBrokerConsumerService::new(kafka_consumer, order_accepted_use_case);

    message_broker_consumer_service.start_consuming();
}
