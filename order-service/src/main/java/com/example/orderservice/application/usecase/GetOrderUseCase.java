package com.example.orderservice.application.usecase;

import com.example.orderservice.application.interfaces.IOrderRepository;
import com.example.orderservice.application.logic.Result;
import com.example.orderservice.domain.entity.Order;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

@Service
public class GetOrderUseCase {
  @Autowired private IOrderRepository orderRepository;

  public Result<Order> getOrderById(Integer id) {
    Result<Order> createdOrderResult = orderRepository.getOrderById(id);
    return createdOrderResult;
  }
}
