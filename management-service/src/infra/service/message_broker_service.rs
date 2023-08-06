use crate::application::interfaces::{
    i_message_broker_consumer::IMessageBrokerConsumer,
    i_message_broker_service::IMessageBrokerService,
};

pub struct MessageBrokerService<'a> {
    consumer: &'a mut dyn IMessageBrokerConsumer,
}

impl<'a> MessageBrokerService<'a> {
    pub fn new(consumer: &'a mut dyn IMessageBrokerConsumer) -> Self {
        Self { consumer }
    }
}

impl<'a> IMessageBrokerService for MessageBrokerService<'a> {
    fn consume_messages(&mut self) -> Vec<String> {
        match self.consumer.consume_messages() {
            Ok(messages) => messages,
            Err(_) => vec![],
        }
    }
}
