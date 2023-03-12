package com.example.orderservice.infra.messagebroker;

import com.example.orderservice.application.usecase.UpdateOrderUseCase;
import com.example.orderservice.domain.enums.EOrderStatus;
import com.example.orderservice.domain.event.OrderEvent;
import com.google.gson.Gson;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.kafka.annotation.KafkaListener;
import org.springframework.stereotype.Component;

@Component
public class KafkaConsumer {
  @Autowired UpdateOrderUseCase updateOrderUseCase;

  private Gson gson = new Gson();
  private OrderEvent event;

  @KafkaListener(topics = "edms", groupId = "edms-group-1")
  public void listenGroupFoo(String message) {
    event = gson.fromJson(message, OrderEvent.class);

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
    updateOrderUseCase.updateOrder(event.getOrderId(), EOrderStatus.COMPLETED,
                                   null);
  }
  public void inventoryNotFound() {
    updateOrderUseCase.updateOrder(event.getOrderId(), EOrderStatus.CANCELLED,
                                   null);
  }
  public void inventoryReserved() {
    updateOrderUseCase.updateOrder(event.getOrderId(), null, true);
  }
  public void transactionCompleted() {
    updateOrderUseCase.updateOrder(event.getOrderId(), EOrderStatus.APPROVED,
                                   null);
  }
  public void transactionFailed() {
    updateOrderUseCase.updateOrder(event.getOrderId(), EOrderStatus.CANCELLED,
                                   null);
  }
}
