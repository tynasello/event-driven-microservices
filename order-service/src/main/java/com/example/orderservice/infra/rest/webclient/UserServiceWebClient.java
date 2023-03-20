package com.example.orderservice.infra.rest.webclient;

import com.example.orderservice.application.interfaces.IUserServiceWebClient;
import com.example.orderservice.application.logic.Result;
import com.example.orderservice.domain.enums.ECustomErrorCode;
import com.example.orderservice.infra.rest.dto.GetUserFromUserServiceResponseDto;
import org.springframework.http.HttpEntity;
import org.springframework.http.HttpHeaders;
import org.springframework.http.HttpMethod;
import org.springframework.http.ResponseEntity;
import org.springframework.stereotype.Component;
import org.springframework.web.client.RestTemplate;

@Component
public class UserServiceWebClient implements IUserServiceWebClient {
  static String UserServiceUrl = "http://user-service:3000";

  public Result<String> getUsernameById(String accessToken) {

    RestTemplate restTemplate = new RestTemplate();

    HttpHeaders headers = new HttpHeaders();
    headers.add("Cookie", "access-token=" + accessToken);

    ResponseEntity<GetUserFromUserServiceResponseDto> response;
    try {
      response = restTemplate.exchange(
          UserServiceWebClient.UserServiceUrl + "/authenticate-user",
          HttpMethod.GET, new HttpEntity<String>(headers),
          GetUserFromUserServiceResponseDto.class);
    } catch (Exception e) {
      return Result.fail(ECustomErrorCode.AUTHENTICATION_ERROR, e.getMessage());
    }

    String username = response.getBody().getUsername();
    if (username == null || username.isEmpty()) {
      return Result.fail(ECustomErrorCode.AUTHENTICATION_ERROR,
                         "Failed to authenticate user");
    }
    return Result.ok(username);
  }
}
