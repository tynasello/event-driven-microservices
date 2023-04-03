use serde::{Deserialize, Serialize};

use crate::domain::enums;

#[derive(Debug, Serialize, Deserialize)]
pub struct OrderShippedEvent {
    #[serde(rename(serialize = "eventType", deserialize = "eventType"))]
    pub event_type: String,
    #[serde(rename(serialize = "orderId", deserialize = "orderId"))]
    pub order_id: i64,
}

impl OrderShippedEvent {
    pub fn new(order_id: i64) -> Self {
        let e_order_accpeted = enums::e_order_event::EOrderEvent::OrderShipped.to_string();
        Self {
            event_type: e_order_accpeted,
            order_id,
        }
    }
}
