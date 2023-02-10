package com.example.orderservice.controller;

import com.example.orderservice.model.OrderModel;
import com.example.orderservice.service.OrderService;
import java.util.List;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping("/orders")
public class OrderController {
  @Autowired private OrderService orderService;

  @GetMapping
  public List<OrderModel> getAllOrders() {
    return orderService.getAllOrders();
  }

  @PostMapping
  public OrderModel addOrder(@RequestBody OrderModel orderRequestDto) {
    return orderService.addOrder(orderRequestDto);
  }
}
