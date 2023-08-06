use crate::{
    application::interfaces::{
        i_message_broker_consumer::MockIMessageBrokerConsumer,
        i_message_broker_service::IMessageBrokerService,
    },
    infra::service::message_broker_service::MessageBrokerService,
};

#[test]
fn test_using_consumer() {
    let mut consumer_mock = MockIMessageBrokerConsumer::new();
    consumer_mock
        .expect_consume_messages()
        .times(1)
        .returning(|| Ok(Vec::new()));

    let mut message_broker_service = MessageBrokerService::new(&mut consumer_mock);
    assert_eq!(message_broker_service.consume_messages().len(), 0);
}

#[test]
fn test_consuming_message() {
    let mut consumer_mock = MockIMessageBrokerConsumer::new();
    consumer_mock
        .expect_consume_messages()
        .times(1)
        .returning(|| Ok(vec!["m1".to_string(), "m2".to_string()]));
    let mut message_broker_service = MessageBrokerService::new(&mut consumer_mock);
    let messages = message_broker_service.consume_messages();
    assert_eq!(messages.len(), 2);
    assert_eq!(messages, vec!["m1", "m2"])
}

#[test]
fn test_error_from_consuming_messages() {
    let mut consumer_mock = MockIMessageBrokerConsumer::new();
    consumer_mock
        .expect_consume_messages()
        .times(1)
        .returning(|| Err("error".to_string()));
    let mut message_broker_service = MessageBrokerService::new(&mut consumer_mock);
    let messages = message_broker_service.consume_messages();
    assert_eq!(messages.len(), 0);
    assert_eq!(messages, Vec::<String>::new())
}
