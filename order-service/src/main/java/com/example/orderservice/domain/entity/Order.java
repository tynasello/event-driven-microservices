package com.example.orderservice.domain.entity;

import com.example.orderservice.domain.enums.EOrderStatus;

public class Order {
  private Integer id;
  private String ordererUsername;
  private EOrderStatus status;
  private String productName;
  private Integer productQuantity;
  private Boolean isInventoryReserved;

  public Order(Integer id, String name, EOrderStatus status, String productName,
               Integer productQuantity, Boolean isInventoryReserved) {
    this.id = id;
    this.ordererUsername = name;
    this.status = status;
    this.productName = productName;
    this.productQuantity = productQuantity;
    this.isInventoryReserved = isInventoryReserved;
  }

  public String toString() {
    return String.format(
        "id: %d, ordererUsername: %s, status: %s, productName: %s, productQuantity: %s, isInventoryReserved: %b",
        id, ordererUsername, status, productName, productQuantity,
        isInventoryReserved);
  }

  public Integer getId() { return id; }

  public String getOrdererUsername() { return ordererUsername; }

  public EOrderStatus getStatus() { return status; }

  public void setStatus(EOrderStatus status) { this.status = status; }

  public String getProductName() { return productName; }

  public Integer getProductQuantity() { return productQuantity; }

  public Boolean getIsInventoryReserved() { return isInventoryReserved; }

  public void setIsInventoryReserved(Boolean isInventoryReserved) {
    this.isInventoryReserved = isInventoryReserved;
  }
}
