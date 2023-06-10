package tests

import (
	"testing"

	"example.com/user-service/src/application/usecase"
	"example.com/user-service/src/domain/entity"
	"example.com/user-service/src/infra/repository/repositorymocks"
	"example.com/user-service/src/infra/service/servicemocks"
)

func TestLoginUseCaseExecute(t *testing.T) {

	userRepositoryMock := &repositorymocks.UserRepositoryMock{}
	hashServiceMock := &servicemocks.HashServiceMock{}
	authTokenServiceMock := &servicemocks.JwtAuthTokenServiceMock{}
	loginUseCase := usecase.LoginUseCase{UserRepository: userRepositoryMock, HashService: hashServiceMock, AuthTokenService: authTokenServiceMock}

	loginWithNonExistingUserResult := loginUseCase.Execute("non-existing-username", "wrong-password")
	if loginWithNonExistingUserResult.IsFailure != true {
		t.Error("Expected loginWithNonExistingUserResult.IsFailure to be true")
	}

	userToCreate := entity.User{
		Id:       1,
		Username: "username",
		Password: "password",
	}
	userRepositoryMock.Create(userToCreate)

	loginWithWrongPasswordResult := loginUseCase.Execute(userToCreate.Username, "wrong-password")
	if loginWithWrongPasswordResult.IsFailure != true {
		t.Error("Expected loginWithWrongPasswordResult.IsFailure to be true")
	}

	loginWithCorrectPasswordResult := loginUseCase.Execute(userToCreate.Username, userToCreate.Password)
	if loginWithCorrectPasswordResult.IsSuccess != true {
		t.Error("Expected loginWithCorrectPasswordResult.IsSuccess to be true")
	}
	accessToken, _ := loginWithCorrectPasswordResult.GetValue()
	if accessToken == nil {
		t.Error("Expected accessToken to not be nil")
	}

}
