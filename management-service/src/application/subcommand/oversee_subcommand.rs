use crate::{
    application::interfaces::i_message_broker_service::IMessageBrokerService,
    infra::{
        message_broker::kafka_consumer::{setup_kafka_consumer, KafkaConsumer},
        service::message_broker_service::MessageBrokerService,
    },
};

pub fn oversee_subcommand() {
    let mut kafka_consumer = setup_kafka_consumer();
    let mut message_broker_consumer = KafkaConsumer::new(&mut kafka_consumer);
    let mut message_broker_service = MessageBrokerService::new(&mut message_broker_consumer);
    loop {
        let messages = message_broker_service.consume_messages();
        for message in messages {
            println!("{}", message);
        }
    }
}
