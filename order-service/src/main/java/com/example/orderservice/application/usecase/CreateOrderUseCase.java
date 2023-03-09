package com.example.orderservice.application.usecase;

import com.example.orderservice.application.interfaces.IOrderRepository;
import com.example.orderservice.application.logic.Result;
import com.example.orderservice.domain.entity.Order;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

@Service
public class CreateOrderUseCase {
  @Autowired private IOrderRepository orderRepository;

  public Result<Order> createOrder(Order order) {

    Result<Order> createdOrderResult = orderRepository.saveOrder(order);
    return createdOrderResult;
  }
}
