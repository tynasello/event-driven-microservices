use crate::{
    application::interfaces::i_message_broker_producer_service::IMessageBrokerProducerService,
    domain::event::order_shipped_event::OrderShippedEvent,
};

pub struct OrderAcceptedUsecase<'a> {
    pub message_broker_producer_service: &'a mut dyn IMessageBrokerProducerService,
}

impl<'a> OrderAcceptedUsecase<'a> {
    pub fn execute(&mut self, order_id: i64) {
        // Some other business logic can occur here

        let order_shipped_event = OrderShippedEvent::new(order_id);
        let order_shipped_event_json = match serde_json::to_string(&order_shipped_event) {
            Ok(order_shipped_event_json) => order_shipped_event_json,
            Err(e) => {
                println!("Error converting order shipped event to json: {}", e);
                return;
            }
        };
        self.message_broker_producer_service
            .publish_message(&order_shipped_event_json);
    }
}
