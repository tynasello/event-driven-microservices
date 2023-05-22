use crate::{
    application::{
        interfaces::i_message_broker_consumer_service::IMessageBrokerConsumerService,
        usecase::order_accepted_usecase::OrderAcceptedUsecase,
    },
    domain::{enums::e_order_event::EOrderEvent, event::order_event::OrderEvent},
    infra::message_broker::kafka_config::KafkaConsumer,
};

pub struct MessageBrokerConsumerService<'a> {
    kafka_consumer: &'a mut KafkaConsumer,
    order_accepted_use_case: &'a mut OrderAcceptedUsecase<'a>,
}

impl<'a> MessageBrokerConsumerService<'a> {
    pub fn new(
        kafka_consumer: &'a mut KafkaConsumer,
        order_accepted_use_case: &'a mut OrderAcceptedUsecase<'a>,
    ) -> Self {
        Self {
            kafka_consumer,
            order_accepted_use_case,
        }
    }

    fn order_accepted(&mut self, order_event: &OrderEvent) {
        self.order_accepted_use_case.execute(order_event.order_id);
    }
}

impl<'a> IMessageBrokerConsumerService for MessageBrokerConsumerService<'a> {
    fn start_consuming(&mut self) {
        let order_event = &mut OrderEvent {
            event_type: "".to_string(),
            order_id: 0,
        };

        loop {
            let messages = self.kafka_consumer.consume_from_kafka();
            if messages.is_empty() {
                continue;
            }

            for message in messages {
                let order_event_result: Result<OrderEvent, serde_json::Error> =
                    serde_json::from_str(&message);

                match order_event_result {
                    Ok(ok_order_event) => {
                        *order_event = ok_order_event;
                    }
                    Err(e) => {
                        println!("Error consuming message: {}", e);
                        continue;
                    }
                }

                if !order_event.is_valid_to_consume() {
                    continue;
                }

                if order_event.event_type == EOrderEvent::OrderAccepted.to_string() {
                    self.order_accepted(order_event);
                }
            }
        }
    }
}
