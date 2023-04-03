package com.example.orderservice.domain.event;

import com.example.orderservice.domain.enums.EOrderEventType;

public class OrderEvent {
  private String eventType;
  private Integer orderId;

  public String toString() {
    return "OrderEvent: " + eventType + " " + orderId;
  }

  public boolean isValidEvent() {
    if (eventType == null) {
      return false;
    }
    try {
      EOrderEventType.valueOf(eventType);
    } catch (IllegalArgumentException e) {
      return false;
    }
    return true;
  }

  public EOrderEventType getEventType() {
    return EOrderEventType.valueOf(eventType);
  }
  public Integer getOrderId() { return orderId; }
}
