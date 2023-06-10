package tests

import (
	"testing"

	"example.com/user-service/src/application/usecase"
	"example.com/user-service/src/domain/entity"
	"example.com/user-service/src/infra/repository/repositorymocks"
	"example.com/user-service/src/infra/service/servicemocks"
)

func TestSignupUseCaseExecute(t *testing.T) {

	userRepositoryMock := &repositorymocks.UserRepositoryMock{}
	hashServiceMock := &servicemocks.HashServiceMock{}
	authTokenServiceMock := &servicemocks.JwtAuthTokenServiceMock{}
	signupUseCase := usecase.SignupUseCase{UserRepository: userRepositoryMock, HashService: hashServiceMock, AuthTokenService: authTokenServiceMock}

	userToCreate := entity.User{
		Id:       1,
		Username: "username",
		Password: "password",
	}

	signupWithNonExistingUserResult := signupUseCase.Execute(userToCreate.Username, userToCreate.Password)
	if signupWithNonExistingUserResult.IsSuccess != true {
		t.Error("Expected loginWithNonExistingUserResult.IsSuccess to be true")
	}
	accessToken, _ := signupWithNonExistingUserResult.GetValue()
	if accessToken == nil {
		t.Error("Expected accessToken to not be nil")
	}
	createdUsersPassword := userRepositoryMock.Users[0].Password
	if createdUsersPassword == userToCreate.Password {
		t.Error("Expected createdUsersPassword to be hashed and not equal userToCreate.Password")
	}

	signupWithExistingUsernameResult := signupUseCase.Execute(userToCreate.Username, "password")
	if signupWithExistingUsernameResult.IsFailure != true {
		t.Error("Expected signupWithExistingUsernameResult.IsFailure to be true")
	}
}
