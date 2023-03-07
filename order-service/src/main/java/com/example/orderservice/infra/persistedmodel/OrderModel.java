package com.example.orderservice.infra.persistedmodel;

import com.example.orderservice.domain.EOrderStatus;
import jakarta.persistence.Entity;
import jakarta.persistence.GeneratedValue;
import jakarta.persistence.GenerationType;
import jakarta.persistence.Id;
import jakarta.persistence.Table;

@Entity
@Table(name = "orders")
public class OrderModel {
  @Id @GeneratedValue(strategy = GenerationType.AUTO) private Long id;

  private String ordererUsername;
  private String status;
  private String productName;
  private String productQuantity;
  private boolean inventoryReserved;

  public OrderModel() {}

  public OrderModel(String name, EOrderStatus status, String productName,
                    String productQuantity, boolean inventoryReserved) {
    this.ordererUsername = name;
    this.status = status.name();
    this.productName = productName;
    this.productQuantity = productQuantity;
    this.inventoryReserved = inventoryReserved;
  }

  public Long getId() { return id; }

  public void setId(Long id) { this.id = id; }

  public String getOrdererUsername() { return ordererUsername; }

  public void setOrdererUsername(String ordererUsername) {
    this.ordererUsername = ordererUsername;
  }

  public EOrderStatus getStatus() { return EOrderStatus.valueOf(status); }

  public void setStatus(EOrderStatus status) { this.status = status.name(); }

  public String getProductName() { return productName; }

  public void setProductName(String productName) {
    this.productName = productName;
  }

  public String getProductQuantity() { return productQuantity; }

  public void setProductQuantity(String productQuantity) {
    this.productQuantity = productQuantity;
  }

  public boolean isInventoryReserved() { return inventoryReserved; }

  public void setInventoryReserved(boolean inventoryReserved) {
    this.inventoryReserved = inventoryReserved;
  }
}
