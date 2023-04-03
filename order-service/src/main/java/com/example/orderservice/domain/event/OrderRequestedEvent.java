package com.example.orderservice.domain.event;

import com.example.orderservice.domain.enums.EOrderEventType;

public class OrderRequestedEvent {
  private String eventType;
  private Integer orderId;
  private String inventoryLabel;
  private Integer inventoryQuantity;

  public OrderRequestedEvent(Integer orderId, String inventoryLabel,
                             Integer inventoryQuantity) {
    this.eventType = EOrderEventType.ORDER_REQUESTED.toString();
    this.orderId = orderId;
    this.inventoryLabel = inventoryLabel;
    this.inventoryQuantity = inventoryQuantity;
  }

  public String toString() {
    return "OrderRequestedEvent: " + orderId + " " + inventoryLabel + " " +
        inventoryQuantity;
  }

  public String getEventType() { return eventType; }

  public Integer getOrderId() { return orderId; }

  public String getInventoryLabel() { return inventoryLabel; }

  public Integer getInventoryQuantity() { return inventoryQuantity; }
}
