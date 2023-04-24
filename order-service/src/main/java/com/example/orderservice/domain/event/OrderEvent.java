package com.example.orderservice.domain.event;

import com.example.orderservice.domain.enums.EOrderEventType;

public class OrderEvent {
  private String eventType;
  private Integer orderId;
  private Boolean isInventoryReserved;
  private Integer inventoryQuantity;
  private String inventoryLabel;

  public String toString(

  ) {
    return "OrderEvent{"
        + "eventType='" + eventType + '\'' + ", orderId=" + orderId +
        ", isInventoryReserved=" + isInventoryReserved +
        ", iventoryQuantity=" + inventoryQuantity + ", inventoryLabel='" +
        inventoryLabel + '\'' + '}';
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
  public Boolean getIsInventoryReserved() { return isInventoryReserved; }
  public Integer getInventoryQuantity() { return inventoryQuantity; }
  public String getInventoryLabel() { return inventoryLabel; }
}
