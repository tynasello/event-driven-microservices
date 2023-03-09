package com.example.orderservice.infra.repository;

import com.example.orderservice.application.interfaces.IOrderRepository;
import com.example.orderservice.application.logic.Result;
import com.example.orderservice.domain.entity.Order;
import com.example.orderservice.domain.enums.ECustomErrorCode;
import com.example.orderservice.domain.enums.EOrderStatus;
import com.example.orderservice.infra.persistedmodel.OrderModel;
import com.example.orderservice.infra.services.jparepository.JpaOrderRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Repository;

@Repository
public class OrderRepository implements IOrderRepository {

  @Autowired private JpaOrderRepository jpaOrderRepository;

  public Result<Order> saveOrder(Order order) {
    OrderModel orderModel = new OrderModel(order);
    OrderModel savedOrderModel;
    try {
      savedOrderModel = jpaOrderRepository.save(orderModel);
    } catch (Error e) {
      return Result.fail(ECustomErrorCode.INTERNAL_SERVER_ERROR,
                         "Order could not be saved");
    }

    Order savedOrder = new Order(
        savedOrderModel.getId(), savedOrderModel.getOrdererUsername(),
        EOrderStatus.valueOf(savedOrderModel.getStatus()),
        savedOrderModel.getProductName(), savedOrderModel.getProductQuantity(),
        savedOrderModel.getIsInventoryReserved());
    return Result.ok(savedOrder);
  }

  public Result<Order> getOrderById(Integer id) {
    OrderModel orderModel;
    try {
      orderModel =
          jpaOrderRepository.findById(Long.valueOf(id.longValue())).get();
    } catch (Error e) {
      return Result.fail(ECustomErrorCode.INTERNAL_SERVER_ERROR,
                         "Order could not be found");
    }

    Order order =
        new Order(orderModel.getId(), orderModel.getOrdererUsername(),
                  EOrderStatus.valueOf(orderModel.getStatus()),
                  orderModel.getProductName(), orderModel.getProductQuantity(),
                  orderModel.getIsInventoryReserved());
    return Result.ok(order);
  }

  public Result<Order> updateOrder(Order order) {
    OrderModel orderModel = new OrderModel(order);
    OrderModel savedOrderModel;
    try {
      savedOrderModel = jpaOrderRepository.save(orderModel);
    } catch (Error e) {
      return Result.fail(ECustomErrorCode.INTERNAL_SERVER_ERROR,
                         "Order could not be updated");
    }

    Order savedOrder = new Order(
        savedOrderModel.getId(), savedOrderModel.getOrdererUsername(),
        EOrderStatus.valueOf(savedOrderModel.getStatus()),
        savedOrderModel.getProductName(), savedOrderModel.getProductQuantity(),
        savedOrderModel.getIsInventoryReserved());
    return Result.ok(savedOrder);
  }
}
