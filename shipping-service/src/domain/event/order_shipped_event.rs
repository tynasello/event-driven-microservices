use serde::{Deserialize, Serialize};

use crate::domain::enums::e_order_event::EOrderEvent;

#[derive(Debug, Serialize, Deserialize)]
pub struct OrderShippedEvent {
    #[serde(rename(serialize = "eventType", deserialize = "eventType"))]
    pub event_type: String,
    #[serde(rename(serialize = "orderId", deserialize = "orderId"))]
    pub order_id: i64,
}

impl OrderShippedEvent {
    pub fn new(order_id: i64) -> Self {
        Self {
            event_type: EOrderEvent::OrderShipped.to_string(),
            order_id,
        }
    }
}
