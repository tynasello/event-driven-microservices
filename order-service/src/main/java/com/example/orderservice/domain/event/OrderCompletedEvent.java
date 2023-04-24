
package com.example.orderservice.domain.event;

import com.example.orderservice.domain.enums.EOrderEventType;

public class OrderCompletedEvent {
  private String eventType;
  private Integer orderId;

  public OrderCompletedEvent(Integer orderId) {
    this.eventType = EOrderEventType.ORDER_COMPLETED.toString();
    this.orderId = orderId;
  }

  public String toString() {
    return "OrderCompletedEvent [eventType=" + eventType +
        ", orderId=" + orderId + "]";
  }

  public String getEventType() { return eventType; }
  public Integer getOrderId() { return orderId; }
}
