package com.example.orderservice.domain.event;

public class OrderRequestedEvent {
  private Integer id;

  public OrderRequestedEvent(Integer id) { this.id = id; }

  public String toString() { return "id=" + this.getId(); }

  public Integer getId() { return id; }

  public void setId(Integer id) { this.id = id; }
}
