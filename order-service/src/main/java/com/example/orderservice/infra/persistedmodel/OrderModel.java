package com.example.orderservice.infra.persistedmodel;

import com.example.orderservice.domain.entity.Order;
import com.example.orderservice.domain.enums.EOrderStatus;
import jakarta.persistence.Entity;
import jakarta.persistence.GeneratedValue;
import jakarta.persistence.GenerationType;
import jakarta.persistence.Id;
import jakarta.persistence.Table;

@Entity
@Table(name = "orders")
public class OrderModel {
  @Id @GeneratedValue(strategy = GenerationType.AUTO) private Integer id;

  private String ordererUsername;
  private String status;
  private String productName;
  private Integer productQuantity;
  private Boolean inventoryReserved;

  public OrderModel() {}

  public OrderModel(Order order) {
    this.id = order.getId();
    this.ordererUsername = order.getOrdererUsername();
    this.status = order.getStatus().name();
    this.productName = order.getProductName();
    this.productQuantity = order.getProductQuantity();
    this.inventoryReserved = order.getIsInventoryReserved();
  }

  public Integer getId() { return id; }

  public void setId(Integer id) { this.id = id; }

  public String getOrdererUsername() { return ordererUsername; }

  public void setOrdererUsername(String ordererUsername) {
    this.ordererUsername = ordererUsername;
  }

  public String getStatus() { return status; }

  public void setStatus(EOrderStatus status) { this.status = status.name(); }

  public String getProductName() { return productName; }

  public void setProductName(String productName) {
    this.productName = productName;
  }

  public Integer getProductQuantity() { return productQuantity; }

  public void setProductQuantity(Integer productQuantity) {
    this.productQuantity = productQuantity;
  }

  public Boolean getIsInventoryReserved() { return inventoryReserved; }

  public void setInventoryReserved(Boolean inventoryReserved) {
    this.inventoryReserved = inventoryReserved;
  }
}
