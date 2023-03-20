package com.example.orderservice.domain.event;

import com.example.orderservice.domain.enums.EOrderEventType;

public class OrderCancelledEvent {
  private String eventType;
  private Integer orderId;
  private Boolean isInventoryReserved;
  private String productName;
  private Integer productQuantity;

  public OrderCancelledEvent(Integer orderId, Boolean isInventoryReserved,
                             String productName, Integer productQuantity) {
    this.eventType = EOrderEventType.ORDER_CANCELLED.toString();
    this.orderId = orderId;
    this.isInventoryReserved = isInventoryReserved;
    this.productName = productName;
    this.productQuantity = productQuantity;
  }

  public String toString() {
    return "OrderCancelledEvent: " + eventType + " " + orderId + " " +
        isInventoryReserved + " " + productName + " " + productQuantity;
  }

  public Integer getOrderId() { return orderId; }

  public String getEventType() { return eventType; }

  public Boolean getIsInventoryReserved() { return isInventoryReserved; }

  public String getProductName() { return productName; }

  public Integer getProductQuantity() { return productQuantity; }
}
