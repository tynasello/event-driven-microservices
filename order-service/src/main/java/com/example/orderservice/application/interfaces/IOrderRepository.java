package com.example.orderservice.application.interfaces;

import com.example.orderservice.application.logic.Result;
import com.example.orderservice.domain.entity.Order;
import org.springframework.stereotype.Repository;

@Repository
public interface IOrderRepository {
  Result<Order> saveOrder(Order order);
  Result<Order> getOrderById(Integer id);
  Result<Order> updateOrder(Order order);
}
