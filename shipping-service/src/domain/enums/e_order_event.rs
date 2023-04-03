pub enum EOrderEvent {
    OrderAccepted,
    OrderShipped,
}

impl EOrderEvent {
    pub fn to_string(&self) -> String {
        match self {
            EOrderEvent::OrderAccepted => "ORDER_ACCEPTED".to_string(),
            EOrderEvent::OrderShipped => "ORDER_SHIPPED".to_string(),
        }
    }
}
