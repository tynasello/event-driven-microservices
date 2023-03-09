package com.example.orderservice.infra.rest.dto;

import com.example.orderservice.domain.entity.Order;

public class OrderResponseDto {
  private Integer id;
  private String ordererUsername;
  private String status;
  private String productName;
  private Integer productQuantity;
  private Boolean isInventoryReserved;

  public OrderResponseDto(Order order) {
    this.id = order.getId();
    this.ordererUsername = order.getOrdererUsername();
    this.status = order.getStatus().toString();
    this.productName = order.getProductName();
    this.productQuantity = order.getProductQuantity();
    this.isInventoryReserved = order.getIsInventoryReserved();
  }

  public String toString() {
    return String.format(
        "id: %d, ordererUsername: %s, status: %s, productName: %s, productQuantity: %s, isInventoryReserved: %b",
        id, ordererUsername, status, productName, productQuantity,
        isInventoryReserved);
  }

  public Integer getId() { return id; }
  public String getOrdererUsername() { return ordererUsername; }
  public String getStatus() { return status; }
  public String getProductName() { return productName; }
  public Integer getProductQuantity() { return productQuantity; }
  public Boolean getIsInventoryReserved() { return isInventoryReserved; }
}
