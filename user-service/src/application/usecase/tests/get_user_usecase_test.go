package tests

import (
	"testing"

	"example.com/user-service/src/application/usecase"
	"example.com/user-service/src/domain/entity"
	"example.com/user-service/src/infra/repository/repositorymocks"
)

func TestGetUserUseCaseExecute(t *testing.T) {

	userRepositoryMock := &repositorymocks.UserRepositoryMock{}
	getUserUseCase := usecase.GetUserUseCase{UserRepository: userRepositoryMock}

	getNonExistingUserResult := getUserUseCase.Execute("non-existing-username")

	if getNonExistingUserResult.IsFailure != true {
		t.Error("Expected getUserResult.IsFailure to be true")
	}

	userToCreate := entity.User{
		Id:       1,
		Username: "username",
		Password: "password",
	}
	userRepositoryMock.Create(userToCreate)

	getExistingUserResult := getUserUseCase.Execute(userToCreate.Username)
	if getExistingUserResult.IsSuccess != true {
		t.Error("Expected getUserResult.IsSuccess to be true")
	}

	existingUser, _ := getExistingUserResult.GetValue()
	if existingUser.Equals(userToCreate) != true {
		t.Error("Expected existingUser to equal userToCreate")
	}

}
