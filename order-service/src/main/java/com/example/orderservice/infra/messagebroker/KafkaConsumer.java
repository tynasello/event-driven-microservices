package com.example.orderservice.infra.messagebroker;

import com.example.orderservice.application.interfaces.IMessageBrokerService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.kafka.annotation.KafkaListener;
import org.springframework.stereotype.Component;

@Component
public class KafkaConsumer {
  @Autowired private IMessageBrokerService messageBrokerService;

  @KafkaListener(topics = "edms", groupId = "order-service-consumer-group")
  public void listenGroupEdms(String message) {
    messageBrokerService.consume(message);
  }
}
