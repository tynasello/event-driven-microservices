use mockall::automock;

#[automock]
pub trait IMessageBrokerConsumer {
    fn consume_messages(&mut self) -> Result<Vec<String>, String>;
}
