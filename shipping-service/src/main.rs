mod enums;

use enums::BrokerEvent;
use kafka::{
    consumer::{Consumer, FetchOffset},
    producer::{Producer, Record},
};

fn main() {
    consume_broker_events()
}

fn get_host() -> Vec<String> {
    vec!["kafka:9092".to_string()]
}

fn get_topic() -> String {
    "edms".to_string()
}

fn consume_broker_events() {
    let hosts = get_host();
    let topic = get_topic();
    let mut consumer = Consumer::from_hosts(hosts)
        .with_topic(topic)
        .with_fallback_offset(FetchOffset::Latest)
        .create()
        .unwrap();

    loop {
        for ms in consumer.poll().unwrap().iter() {
            for m in ms.messages() {
                let message = std::str::from_utf8(m.value).unwrap();
                if message == "OrderRequestAccepted" {
                    handle_broker_event(BrokerEvent::OrderRequestAccepted);
                }
            }

            consumer.consume_messageset(ms).unwrap();
        }

        consumer.commit_consumed().unwrap();
    }
}

fn produce_broker_events() {
    let hosts = get_host();
    let topic = get_topic();
    let mut producer = Producer::from_hosts(hosts).create().unwrap();

    for i in 0..10 {
        let buf = format!("{i}");
        producer
            .send(&Record::from_value(&topic, buf.as_bytes()))
            .unwrap();
        println!("Sent: {i}");
    }
}

fn handle_broker_event(broker_event: BrokerEvent) {
    if let BrokerEvent::OrderRequestAccepted = broker_event {
        println!("Inventory found");
    }
}
