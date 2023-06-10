package logic

import "errors"

type Result[T any] struct {
	IsSuccess     bool
	IsFailure     bool
	_value        T
	_errorMessage string
}

func OkResult[T any](value T) *Result[T] {
	return &Result[T]{true, false, value, ""}
}

func FailedResult[T any](errorMessage string) *Result[T] {
	var t T
	return &Result[T]{false, true, t, errorMessage}
}

func (r Result[T]) GetValue() (*T, error) {
	if r.IsFailure {
		return nil, errors.New("Cannot get value from a failed result")
	}
	return &r._value, nil
}

func (r Result[T]) GetErrorMessage() (*string, error) {
	if r.IsSuccess {
		return nil, errors.New("Cannot get error message from a successful result")
	}
	return &r._errorMessage, nil
}
