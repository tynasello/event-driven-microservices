package com.example.orderservice.infra.services;

import com.example.orderservice.application.interfaces.IJsonService;
import com.example.orderservice.application.interfaces.IMessageBrokerService;
import com.example.orderservice.application.logic.Result;
import com.example.orderservice.application.usecase.UpdateOrderUseCase;
import com.example.orderservice.domain.entity.Order;
import com.example.orderservice.domain.enums.EOrderEventType;
import com.example.orderservice.domain.enums.EOrderStatus;
import com.example.orderservice.domain.event.OrderAcceptedEvent;
import com.example.orderservice.domain.event.OrderCancelledEvent;
import com.example.orderservice.domain.event.OrderEvent;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.kafka.core.KafkaTemplate;

public class MessageBrokerService implements IMessageBrokerService {

  @Autowired UpdateOrderUseCase updateOrderUseCase;
  @Autowired private IJsonService jsonService;
  @Autowired private KafkaTemplate<String, String> kafkaTemplate;

  private OrderEvent event;

  public void send(String topicName, String message) {
    kafkaTemplate.send(topicName, message);
  }

  public void consume(String message) {
    event = jsonService.fromJson(message, OrderEvent.class);

    if (event == null || !event.isValidEvent()) {
      return;
    }

    switch (event.getEventType()) {
    case ORDER_SHIPPED:
      orderShipped();
      break;
    case INVENTORY_NOT_FOUND:
      inventoryNotFound();
      break;
    case INVENTORY_RESERVED:
      inventoryReserved();
      break;
    case TRANSACTION_COMPLETED:
      transactionCompleted();
      break;
    case TRANSACTION_FAILED:
      transactionFailed();
      break;
    default:
      // event is not of concern to order service
      return;
    }
  }

  public void orderShipped() {
    Result<Order> updatedOrderReuslt = updateOrderUseCase.updateOrder(
        event.getOrderId(), EOrderStatus.COMPLETED, null);
    if (updatedOrderReuslt.isFailure) {
      return;
    }
    this.send("edms", EOrderEventType.ORDER_COMPLETED.toString());
  }

  public void inventoryNotFound() {
    Result<Order> updatedOrderReuslt = updateOrderUseCase.updateOrder(
        event.getOrderId(), EOrderStatus.CANCELLED, null);
    if (updatedOrderReuslt.isFailure) {
      return;
    }
    OrderCancelledEvent orderCancelledEvent = new OrderCancelledEvent(
        event.getOrderId(), false, event.getProductName(),
        event.getProductQuantity());
    this.send("edms", jsonService.toJson(orderCancelledEvent));
  }

  public void inventoryReserved() {
    updateOrderUseCase.updateOrder(event.getOrderId(), null, true);
  }

  public void transactionCompleted() {
    Result<Order> updatedOrderReuslt = updateOrderUseCase.updateOrder(
        event.getOrderId(), EOrderStatus.APPROVED, null);
    if (updatedOrderReuslt.isFailure) {
      return;
    }
    OrderAcceptedEvent orderAcceptedEvent =
        new OrderAcceptedEvent(event.getOrderId());
    this.send("edms", jsonService.toJson(orderAcceptedEvent));
  }

  public void transactionFailed() {
    Result<Order> updatedOrderReuslt = updateOrderUseCase.updateOrder(
        event.getOrderId(), EOrderStatus.CANCELLED, null);

    if (updatedOrderReuslt.isFailure) {
      return;
    }
    OrderCancelledEvent orderCancelledEvent = new OrderCancelledEvent(
        event.getOrderId(), event.getIsInventoryReserved(),
        event.getProductName(), event.getProductQuantity());
    this.send("edms", jsonService.toJson(orderCancelledEvent));
  }
}