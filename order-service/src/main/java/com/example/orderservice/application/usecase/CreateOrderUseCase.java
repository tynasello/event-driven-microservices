package com.example.orderservice.application.usecase;

import com.example.orderservice.application.logic.Result;
import com.example.orderservice.infra.persistedmodel.OrderModel;
import com.example.orderservice.infra.repository.OrderRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

@Service
public class CreateOrderUseCase {
  @Autowired private OrderRepository orderRepository;

  public Result<OrderModel> createOrder(OrderModel order) {

    OrderModel createdOrder = orderRepository.save(order);
    return Result.ok(createdOrder);
  }
}
