package com.example.orderservice.application.usecase;

import com.example.orderservice.application.interfaces.IOrderRepository;
import com.example.orderservice.application.logic.Result;
import com.example.orderservice.domain.entity.Order;
import com.example.orderservice.domain.enums.ECustomErrorCode;
import com.example.orderservice.domain.enums.EOrderStatus;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

@Service
public class UpdateOrderUseCase {
  @Autowired private IOrderRepository orderRepository;

  public Result<Order> updateOrder(Integer id, EOrderStatus status,
                                   Boolean isInventoryReserved) {
    Result<Order> existingOrderResult = orderRepository.getOrderById(id);
    if (existingOrderResult.isFailure) {
      return Result.fail(ECustomErrorCode.USER_INPUT_ERROR,
                         "Order does not exits");
    }
    Order existingOrder = existingOrderResult.getValue();
    if (status != null) {
      existingOrder.setStatus(status);
    }
    if (isInventoryReserved != null) {
      existingOrder.setIsInventoryReserved(isInventoryReserved);
    }
    Result<Order> updatedOrder = orderRepository.updateOrder(existingOrder);
    return updatedOrder;
  }
}
