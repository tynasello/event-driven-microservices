package com.example.orderservice.infra.rest.controller;

import com.example.orderservice.application.logic.Result;
import com.example.orderservice.application.usecase.CreateOrderUseCase;
import com.example.orderservice.application.usecase.GetOrderUseCase;
import com.example.orderservice.application.usecase.UpdateOrderUseCase;
import com.example.orderservice.domain.entity.Order;
import com.example.orderservice.domain.enums.ECustomErrorCode;
import com.example.orderservice.domain.enums.EOrderStatus;
import com.example.orderservice.infra.rest.dto.CreateOrderRequestDto;
import com.example.orderservice.infra.rest.dto.GetOrderRequestDto;
import com.example.orderservice.infra.rest.dto.OrderResponseDto;
import com.example.orderservice.infra.rest.dto.UpdateOrderRequestDto;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.web.bind.annotation.ExceptionHandler;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.ResponseStatus;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.server.ResponseStatusException;

@RestController
@RequestMapping()
public class OrderController {
  @Autowired private CreateOrderUseCase createOrderUseCase;
  @Autowired private GetOrderUseCase getOrderUseCase;
  @Autowired private UpdateOrderUseCase updateOrderUseCase;

  public static void HandleFailedResult(Result result) {
    switch (result.getErrorCode()) {
    case USER_INPUT_ERROR:
      throw new ResponseStatusException(HttpStatus.BAD_REQUEST,
                                        result.toString());
    default:
      throw new ResponseStatusException(HttpStatus.INTERNAL_SERVER_ERROR,
                                        result.toString());
    }
  }

  @ExceptionHandler()
  public void handleException(RuntimeException e) {
    OrderController.HandleFailedResult(
        Result.fail(ECustomErrorCode.INTERNAL_SERVER_ERROR, e.getMessage()));
  }

  @GetMapping("get-order")
  @ResponseStatus(code = HttpStatus.OK)
  public OrderResponseDto getOrder(@RequestBody GetOrderRequestDto dto) {

    Result<GetOrderRequestDto> isDtoValidResult =
        GetOrderRequestDto.isValid(dto);
    if (isDtoValidResult.isFailure) {
      OrderController.HandleFailedResult(isDtoValidResult);
    }
    GetOrderRequestDto validDto = isDtoValidResult.getValue();

    Result<Order> orderResult = getOrderUseCase.getOrderById(validDto.getId());

    if (orderResult.isFailure) {
      OrderController.HandleFailedResult(orderResult);
    }
    Order order = orderResult.getValue();

    OrderResponseDto responseDto = new OrderResponseDto(order);
    return responseDto;
  }

  @PostMapping("create-order")
  @ResponseStatus(code = HttpStatus.CREATED)
  public OrderResponseDto createOrder(@RequestBody CreateOrderRequestDto dto) {
    Result<CreateOrderRequestDto> isDtoValidResult =
        CreateOrderRequestDto.isValid(dto);
    if (isDtoValidResult.isFailure) {
      OrderController.HandleFailedResult(isDtoValidResult);
    }
    CreateOrderRequestDto validDto = isDtoValidResult.getValue();
    Order order = new Order(null, validDto.getOrdererUsername(),
                            EOrderStatus.REQUESTED, validDto.getProductName(),
                            validDto.getProductQuantity(), false);
    Result<Order> createdOrderResult = createOrderUseCase.createOrder(order);
    if (createdOrderResult.isFailure) {
      OrderController.HandleFailedResult(createdOrderResult);
    }
    Order createdOrder = createdOrderResult.getValue();

    OrderResponseDto responseDto = new OrderResponseDto(createdOrder);
    return responseDto;
  }

  @PostMapping("update-order")
  @ResponseStatus(code = HttpStatus.CREATED)
  public OrderResponseDto updateOrder(@RequestBody UpdateOrderRequestDto dto) {

    Result<UpdateOrderRequestDto> isDtoValidResult =
        UpdateOrderRequestDto.isValid(dto);
    if (isDtoValidResult.isFailure) {
      OrderController.HandleFailedResult(isDtoValidResult);
    }
    UpdateOrderRequestDto validDto = isDtoValidResult.getValue();

    Result<Order> updatedOrder = updateOrderUseCase.updateOrder(
        validDto.getId(),
        validDto.getStatus() == null
            ? null
            : EOrderStatus.valueOf(validDto.getStatus()),
        validDto.getIsInventoryReserved());
    if (updatedOrder.isFailure) {
      OrderController.HandleFailedResult(updatedOrder);
    }
    Order updatedOrderValue = updatedOrder.getValue();

    OrderResponseDto responseDto = new OrderResponseDto(updatedOrderValue);
    return responseDto;
  }
}
