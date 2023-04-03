package com.example.orderservice.domain.event;

import com.example.orderservice.domain.enums.EOrderEventType;

public class OrderCancelledEvent {
  private String eventType;
  private Integer orderId;
  private Boolean isInventoryReserved;
  private String inventoryLabel;
  private Integer inventoryQuantity;

  public OrderCancelledEvent(Integer orderId, Boolean isInventoryReserved,
                             String inventoryLabel, Integer inventoryQuantity) {
    this.eventType = EOrderEventType.ORDER_CANCELLED.toString();
    this.orderId = orderId;
    this.isInventoryReserved = isInventoryReserved;
    this.inventoryLabel = inventoryLabel;
    this.inventoryQuantity = inventoryQuantity;
  }

  public String toString() {
    return "OrderCancelledEvent: " + eventType + " " + orderId + " " +
        isInventoryReserved + " " + inventoryLabel + " " + inventoryQuantity;
  }

  public Integer getOrderId() { return orderId; }

  public String getEventType() { return eventType; }

  public Boolean getIsInventoryReserved() { return isInventoryReserved; }

  public String getInventoryLabel() { return inventoryLabel; }

  public Integer getInventoryQuantity() { return inventoryQuantity; }
}
