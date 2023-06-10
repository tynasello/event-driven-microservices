package tests

import (
	"testing"

	"example.com/user-service/src/application/logic"
)

func TestResult(t *testing.T) {
	okResult := logic.OkResult[string]("a")
	if okResult.IsSuccess != true {
		t.Error("Expected result.IsSuccess to be true")
	}
	if okResult.IsFailure != false {
		t.Error("Expected result.IsFailure to be false")
	}

	failedResult := logic.FailedResult[string]("error")
	if failedResult.IsSuccess != false {
		t.Error("Expected result.IsSuccess to be false")
	}
	if failedResult.IsFailure != true {
		t.Error("Expected result.IsFailure to be true")
	}

	okResultValue, okResultGetValueError := okResult.GetValue()
	if okResultGetValueError != nil {
		t.Error("Expected okResult.GetValue() to return a value")
	}
	if *okResultValue != "a" {
		t.Error("Expected okResult.GetValue() to return a value")
	}

	failedResultValue, failedResultGetValueError := failedResult.GetValue()
	if failedResultValue != nil {
		t.Error("Expected failedResult.GetValue() to return nil")
	}
	if failedResultGetValueError == nil {
		t.Error("Expected failedResult.GetValue() to return an error")
	}

	okResultErrorMessage, okResultGetErrorMessageError := okResult.GetErrorMessage()
	if okResultErrorMessage != nil {
		t.Error("Expected okResult.GetErrorMessage() to return an empty string")
	}
	if okResultGetErrorMessageError == nil {
		t.Error("Expected okResult.GetErrorMessage() to return an error")
	}

	failedResultErrorMessage, failedResultGetErrorMessageError := failedResult.GetErrorMessage()
	if failedResultErrorMessage == nil {
		t.Error("Expected failedResult.GetErrorMessage() to return a value")
	}
	if failedResultGetErrorMessageError != nil {
		t.Error("Expected failedResult.GetErrorMessage() to return a value")
	}

}
