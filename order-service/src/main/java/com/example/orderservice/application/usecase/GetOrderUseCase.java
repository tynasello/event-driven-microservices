package com.example.orderservice.application.usecase;

import com.example.orderservice.application.logic.Result;
import com.example.orderservice.infra.persistedmodel.OrderModel;
import com.example.orderservice.infra.repository.OrderRepository;
import java.util.List;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

@Service
public class GetOrderUseCase {
  @Autowired private OrderRepository orderRepository;

  public Result<List<OrderModel>> getAllOrders() {
    return Result.ok(orderRepository.findAll());
  }
}
