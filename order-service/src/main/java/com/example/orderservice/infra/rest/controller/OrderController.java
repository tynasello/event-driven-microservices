package com.example.orderservice.infra.rest.controller;

import com.example.orderservice.application.logic.Result;
import com.example.orderservice.application.usecase.CreateOrderUseCase;
import com.example.orderservice.application.usecase.GetOrderUseCase;
import com.example.orderservice.application.usecase.UpdateOrderUseCase;
import com.example.orderservice.infra.persistedmodel.OrderModel;
import java.util.List;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

/**
 *
 */
@RestController
@RequestMapping()
public class OrderController {
  @Autowired private CreateOrderUseCase createOrderUseCase;
  @Autowired private GetOrderUseCase getOrderUseCase;
  @Autowired private UpdateOrderUseCase updateOrderUseCase;

  @GetMapping("get-orders")
  public List<OrderModel> getAllOrders() {
    Result<List<OrderModel>> allOrders = getOrderUseCase.getAllOrders();

    if (allOrders.isFailure) {
      throw new RuntimeException(allOrders.getErrorMessage());
    }

    return allOrders.getValue();
  }

  @PostMapping("create-order")
  public OrderModel createOrder(@RequestBody OrderModel orderRequestDto,
                                OrderModel orderModel) {
    Result<OrderModel> createdOrderModel =
        createOrderUseCase.createOrder(orderModel);
    if (createdOrderModel.isFailure) {
      throw new RuntimeException(createdOrderModel.getErrorMessage());
    }
    return createdOrderModel.getValue();
  }

  @PostMapping("update-order")
  public OrderModel updateOrder(@RequestBody OrderModel orderRequestDto,
                                OrderModel orderModel) {
    Result<OrderModel> updatedOrderModel =
        updateOrderUseCase.updateOrder(orderModel);
    if (updatedOrderModel.isFailure) {
      throw new RuntimeException(updatedOrderModel.getErrorMessage());
    }
    return updatedOrderModel.getValue();
  }
}
