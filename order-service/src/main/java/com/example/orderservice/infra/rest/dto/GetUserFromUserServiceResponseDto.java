package com.example.orderservice.infra.rest.dto;

public class GetUserFromUserServiceResponseDto {
  private String username;
  public GetUserFromUserServiceResponseDto() {}
  public GetUserFromUserServiceResponseDto(String username) {
    this.username = username;
  }
  public String toString() {
    return "GetUserFromUserServiceResponseDto(username=" + this.getUsername() +
        ")";
  }

  public String getUsername() { return username; }

  public void setUsername(String username) { this.username = username; }
}
