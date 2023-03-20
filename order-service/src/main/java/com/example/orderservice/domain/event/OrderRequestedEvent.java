package com.example.orderservice.domain.event;

public class OrderRequestedEvent {
  private Integer orderId;
  private String productName;
  private Integer productQuantity;

  public OrderRequestedEvent(Integer orderId, String productName,
                             Integer productQuantity) {
    this.orderId = orderId;
    this.productName = productName;
    this.productQuantity = productQuantity;
  }

  public String toString() {
    return "OrderRequestedEvent: " + orderId + " " + productName + " " +
        productQuantity;
  }

  public Integer getOrderId() { return orderId; }

  public String getProductName() { return productName; }

  public Integer getProductQuantity() { return productQuantity; }
}
