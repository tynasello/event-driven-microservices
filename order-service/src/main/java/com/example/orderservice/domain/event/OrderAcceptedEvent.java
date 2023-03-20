package com.example.orderservice.domain.event;

import com.example.orderservice.domain.enums.EOrderEventType;

public class OrderAcceptedEvent {
  private String eventType;
  private Integer orderId;

  public OrderAcceptedEvent(Integer orderId) {
    this.eventType = EOrderEventType.ORDER_ACCEPTED.toString();
    this.orderId = orderId;
  }

  public String toString() {
    return "OrderAcceptedEvent: " + eventType + " " + orderId;
  }

  public Integer getOrderId() { return orderId; }

  public String getEventType() { return eventType; }
}
