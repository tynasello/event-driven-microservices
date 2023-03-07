package com.example.orderservice.application.logic;

public class Result<T> {
  public final boolean isSuccess;
  public final boolean isFailure;
  private final Error _errorCode;
  private final String _errorMessage;
  private final T value;

  public Result(boolean isSuccess, T value, Error errorCode,
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
    this.value = value;
    this._errorCode = errorCode;
    this._errorMessage = errorMessage;
  }

  public static <T> Result<T> ok(T value) {
    return new Result<T>(true, value, null, null);
  }

  public static <T> Result<T> fail(Error errorCode, String errorMessage) {
    return new Result<T>(false, null, errorCode, errorMessage);
  }

  public Error getErrorCode() {
    if (this.isSuccess) {
      throw new IllegalStateException("InvalidOperation: Can't get the error "
                                      + "code from a successful result");
    }
    return _errorCode;
  }

  public String getErrorMessage() {
    if (this.isSuccess) {
      throw new IllegalStateException("InvalidOperation: Can't get the error "
                                      + "message from a successful result");
    }
    return _errorMessage;
  }

  public T getValue() {
    if (this.isFailure) {
      throw new IllegalStateException("InvalidOperation: Can't get the value "
                                      + "from a failed result");
    }
    return value;
  }
}
