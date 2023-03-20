package com.example.orderservice.application.interfaces;

public interface IJsonService {
  String toJson(Object object);
  <T> T fromJson(String string, Class<T> classType);
}
