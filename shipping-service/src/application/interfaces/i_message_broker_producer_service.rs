pub trait IMessageBrokerProducerService {
    fn publish_message(&mut self, message: &str);
}
