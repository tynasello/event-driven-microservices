package com.example.orderservice.application.usecase;

import com.example.orderservice.application.logic.Result;
import com.example.orderservice.infra.persistedmodel.OrderModel;
import com.example.orderservice.infra.repository.OrderRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

@Service
public class UpdateOrderUseCase {
  @Autowired private OrderRepository orderRepository;

  public Result<OrderModel> updateOrder(OrderModel order) {
    boolean orderExists = orderRepository.existsById(order.getId());
    if (!orderExists) {
      return Result.fail(null, "Order not found");
    }
    OrderModel updatedOrder = orderRepository.save(order);
    return Result.ok(updatedOrder);
  }
}
