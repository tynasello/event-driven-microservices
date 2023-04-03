use crate::{
    application::{
        interfaces::i_message_broker_consumer_service::IMessageBrokerConsumerService,
        usecase::order_accepted_usecase::OrderAcceptedUsecase,
    },
    domain::{enums, event::order_event::OrderEvent},
    infra::message_broker::kafka_config::KafkaConsumer,
};

pub struct MessageBrokerConsumerService<'a> {
    pub order_event: Box<OrderEvent>,
    pub kafka_consumer: &'a mut KafkaConsumer,
    pub order_accepted_use_case: &'a mut OrderAcceptedUsecase<'a>,
}

impl<'a> IMessageBrokerConsumerService for MessageBrokerConsumerService<'a> {
    fn start_consuming(&mut self) {
        loop {
            let messages = self.kafka_consumer.consume_from_kafka();
            if messages.is_empty() {
                continue;
            }

            for message in messages {
                let order_event_result: Result<OrderEvent, serde_json::Error> =
                    serde_json::from_str(&message);

                match order_event_result {
                    Ok(order_event) => {
                        self.order_event = Box::new(order_event);
                    }
                    Err(e) => {
                        println!("Error consuming message: {}", e);
                        continue;
                    }
                }

                if !self.order_event.is_valid_to_consume() {
                    continue;
                }

                let e_order_accpeted = enums::e_order_event::EOrderEvent::OrderAccepted.to_string();

                if self.order_event.event_type == e_order_accpeted {
                    self.order_accepted();
                }
            }
        }
    }
}
impl<'a> MessageBrokerConsumerService<'a> {
    fn order_accepted(&mut self) {
        self.order_accepted_use_case
            .execute(self.order_event.order_id);
    }
}
