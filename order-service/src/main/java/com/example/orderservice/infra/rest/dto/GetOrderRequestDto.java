package com.example.orderservice.infra.rest.dto;

import com.example.orderservice.application.logic.Result;
import com.example.orderservice.domain.enums.ECustomErrorCode;

public class GetOrderRequestDto {
  private Integer id;

  public GetOrderRequestDto() {}

  public GetOrderRequestDto(Integer id) { this.id = id; }

  public String toString() {
    return "GetOrderRequestDto{"
        + "id=" + id + '}';
  }

  public static Result<GetOrderRequestDto> isValid(GetOrderRequestDto dto) {
    if (dto.id == null) {
      return Result.fail(ECustomErrorCode.USER_INPUT_ERROR, "*id* is required");
    }
    return Result.ok(dto);
  }

  public Integer getId() { return id; }

  public void setId(Integer id) { this.id = id; }
}
