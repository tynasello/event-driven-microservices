use serde::{Deserialize, Serialize};

use crate::domain::enums;

#[derive(Debug, Serialize, Deserialize)]
pub struct OrderEvent {
    #[serde(rename(serialize = "eventType", deserialize = "eventType"))]
    pub event_type: String,
    #[serde(rename(serialize = "orderId", deserialize = "orderId"))]
    pub order_id: i64,
}

impl OrderEvent {
    pub fn is_valid_to_consume(&self) -> bool {
        let order_accepted_event = enums::e_order_event::EOrderEvent::to_string(
            &enums::e_order_event::EOrderEvent::OrderAccepted,
        );
        if self.event_type == order_accepted_event {
            return true;
        }
        return false;
    }
}
