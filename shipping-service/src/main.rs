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
use domain::event::order_event::OrderEvent;

mod application;
mod domain;
mod infra;

fn main() {
    dotenv().ok();

    let mut kafka_producer = KafkaProducer::new();
    let mut kafka_consumer = KafkaConsumer::new();

    let mut message_broker_producer_service = MessageBrokerProducerService {
        kafka_producer: &mut kafka_producer,
    };

    let mut order_accepted_use_case = OrderAcceptedUsecase {
        message_broker_producer_service: &mut message_broker_producer_service,
    };

    let mut message_broker_consumer_service = MessageBrokerConsumerService {
        order_event: Box::new(OrderEvent {
            event_type: "".to_string(),
            order_id: 0,
        }),
        kafka_consumer: &mut kafka_consumer,
        order_accepted_use_case: &mut order_accepted_use_case,
    };

    message_broker_consumer_service.start_consuming();
}
