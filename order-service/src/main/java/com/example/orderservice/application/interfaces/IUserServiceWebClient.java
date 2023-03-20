package com.example.orderservice.application.interfaces;

import com.example.orderservice.application.logic.Result;

public interface IUserServiceWebClient {
  Result<String> getUsernameById(String accessToken);
}
