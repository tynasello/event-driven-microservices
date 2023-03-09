package com.example.orderservice;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.web.bind.annotation.RestController;

@SpringBootApplication()
@RestController
public class OrderServiceApplication {
  public static void main(String[] args) {

    SpringApplication.run(OrderServiceApplication.class, args);

    System.out.println("Order service is running...");
  }
}
