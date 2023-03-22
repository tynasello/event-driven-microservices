package com.example.orderservice.application.logic;

import com.example.orderservice.domain.enums.ECustomErrorCode;

public class Result<T> {
  public final boolean isSuccess;
  public final boolean isFailure;
  private final ECustomErrorCode _errorCode;
  private final String _errorMessage;
  private final T _value;

  public Result(boolean isSuccess, T value, ECustomErrorCode errorCode,
                String errorMessage) {
    if (isSuccess && errorCode != null) {
      throw new IllegalArgumentException("InvalidOperation: A result cannot be "
                                         + "successful and contain an error");
    }
    if (!isSuccess && errorCode == null) {
      throw new IllegalArgumentException("InvalidOperation: A failing result "
                                         + "needs to contain an error message");
    }
    this.isSuccess = isSuccess;
    this.isFailure = !isSuccess;
    this._value = value;
    this._errorCode = errorCode;
    this._errorMessage = errorMessage;
  }

  public String toString() {
    if (this.isSuccess) {
      return this._value.toString();
    } else {
      return this._errorCode + " - " + this._errorMessage;
    }
  }

  public static <T> Result<T> ok(T value) {
    return new Result<T>(true, value, null, null);
  }

  public static <T> Result<T> fail(ECustomErrorCode errorCode,
                                   String errorMessage) {
    return new Result<T>(false, null, errorCode, errorMessage);
  }

  public ECustomErrorCode getErrorCode() {
    if (this.isSuccess) {
      throw new IllegalStateException("InvalidOperation: Can't get the error "
                                      + "code from a successful result");
    }
    return this._errorCode;
  }

  public String getErrorMessage() {
    if (this.isSuccess) {
      throw new IllegalStateException("InvalidOperation: Can't get the error "
                                      + "message from a successful result");
    }
    return this._errorMessage;
  }

  public T get_value() {
    if (this.isFailure) {
      throw new IllegalStateException("InvalidOperation: Can't get the value "
                                      + "from a failed result");
    }
    return this._value;
  }
}
