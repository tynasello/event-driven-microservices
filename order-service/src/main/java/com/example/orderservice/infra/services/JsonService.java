package com.example.orderservice.infra.services;

import com.example.orderservice.application.interfaces.IJsonService;
import com.google.gson.Gson;

public class JsonService implements IJsonService {
  private Gson gson = new Gson();

  public String toJson(Object object) { return gson.toJson(object); }

  public <T> T fromJson(String string, Class<T> classType) {
    return gson.fromJson(string, classType);
  }
}
