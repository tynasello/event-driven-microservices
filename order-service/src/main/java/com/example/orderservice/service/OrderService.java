package com.example.orderservice.service;

import com.example.orderservice.model.OrderModel;
import com.example.orderservice.repository.OrderRepository;
import java.util.List;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

@Service
public class OrderService {
  @Autowired private OrderRepository orderRepository;

  public List<OrderModel> getAllOrders() { return orderRepository.findAll(); }

  public OrderModel addOrder(OrderModel order) {
    return orderRepository.save(order);
  }
}
