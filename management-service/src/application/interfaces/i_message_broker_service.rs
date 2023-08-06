use mockall::automock;

#[automock]
pub trait IMessageBrokerService {
    fn consume_messages(&mut self) -> Vec<String>;
}
