package com.example.orderservice.infra.rest.dto;

import com.example.orderservice.application.logic.Result;
import com.example.orderservice.domain.enums.ECustomErrorCode;

public class CreateOrderRequestDto {
  private String ordererUsername;
  private String productName;
  private Integer productQuantity;

  public CreateOrderRequestDto(String name, String productName,
                               Integer productQuantity) {
    this.ordererUsername = name;
    this.productName = productName;
    this.productQuantity = productQuantity;
  }

  public String toString() {
    return "CreateOrderRequestDto{"
        + "ordererUsername='" + ordererUsername + '\'' + ", productName='" +
        productName + '\'' + ", productQuantity='" + productQuantity + '\'' +
        '}';
  }

  public static Result<CreateOrderRequestDto>
  isValid(CreateOrderRequestDto dto) {
    if (dto.ordererUsername == null || dto.ordererUsername.isEmpty()) {
      return Result.fail(ECustomErrorCode.USER_INPUT_ERROR,
                         "*orderUsername* is required");
    }
    if (dto.productName == null || dto.productName.isEmpty()) {
      return Result.fail(ECustomErrorCode.USER_INPUT_ERROR,
                         "*productName* is required");
    }
    if (dto.productQuantity == null || dto.productQuantity <= 0) {
      return Result.fail(ECustomErrorCode.USER_INPUT_ERROR,
                         "*productQuantity* is invalid");
    }
    return Result.ok(dto);
  }

  public String getOrdererUsername() { return ordererUsername; }

  public String getProductName() { return productName; }

  public Integer getProductQuantity() { return productQuantity; }
}
