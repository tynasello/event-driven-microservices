
package com.example.orderservice.infra.rest.dto;

import com.example.orderservice.application.logic.Result;
import com.example.orderservice.domain.enums.ECustomErrorCode;
import com.example.orderservice.domain.enums.EOrderStatus;

public class UpdateOrderRequestDto {
  private Integer id;
  private String status;
  private Boolean isInventoryReserved;

  public UpdateOrderRequestDto(Integer id, String status,
                               Boolean isInventoryReserved) {
    this.id = id;
    this.status = status;
    this.isInventoryReserved = isInventoryReserved;
  }

  public String toString() {
    return "UpdateOrderRequestDto{"
        + "id=" + id + ", status='" + status + '\'' +
        ", isInventoryReserved=" + isInventoryReserved + '}';
  }

  public static Result<UpdateOrderRequestDto>
  isValid(UpdateOrderRequestDto dto) {
    if (dto.getId() == null) {
      return Result.fail(ECustomErrorCode.USER_INPUT_ERROR, "*id* is required");
    }
    if (dto.getStatus() != null) {
      try {
        EOrderStatus.valueOf(dto.getStatus());
      } catch (IllegalArgumentException e) {
        return Result.fail(ECustomErrorCode.USER_INPUT_ERROR,
                           "*status* is invalid");
      }
      return Result.ok(dto);
    }
    return Result.ok(dto);
  }

  public Integer getId() { return id; }
  public String getStatus() { return status; }
  public Boolean getIsInventoryReserved() { return isInventoryReserved; }
}
