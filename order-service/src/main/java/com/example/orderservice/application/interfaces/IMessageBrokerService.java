package com.example.orderservice.application.interfaces;

public interface IMessageBrokerService {
  void send(String topicName, String message);
  void consume(String message);
}
