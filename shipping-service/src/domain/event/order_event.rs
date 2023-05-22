use serde::{Deserialize, Serialize};

use crate::domain::enums::e_order_event::EOrderEvent;

#[derive(Debug, Serialize, Deserialize)]
pub struct OrderEvent {
    #[serde(rename(serialize = "eventType", deserialize = "eventType"))]
    pub event_type: String,
    #[serde(rename(serialize = "orderId", deserialize = "orderId"))]
    pub order_id: i64,
}

impl OrderEvent {
    pub fn is_valid_to_consume(&self) -> bool {
        if self.event_type == EOrderEvent::to_string(&EOrderEvent::OrderAccepted) {
            return true;
        };
        return false;
    }
}
