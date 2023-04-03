package usecase

import (
	"time"

	"example.com/user-service/src/application/interfaces"
	"example.com/user-service/src/application/logic"
	"example.com/user-service/src/domain/entity"
)

type SignupUseCase struct {
	UserRepository   interfaces.IUserRepository
	HashService      interfaces.IHashService
	AuthTokenService interfaces.IAuthTokenService
}

func (u SignupUseCase) Execute(username string, password string) *logic.Result[string] {
	existingUserResult := u.UserRepository.GetByUsername(username)

	if existingUserResult.IsSuccess {
		return logic.FailedResult[string]("User already exists")
	}

	hashedPassword, err := u.HashService.Hash(password)
	if err != nil {
		return logic.FailedResult[string]("Error hashing user password")
	}

	user := entity.User{
		Username: username,
		Password: hashedPassword,
	}

	createdUserResult := u.UserRepository.Create(user)

	if createdUserResult.IsFailure {
		return logic.FailedResult[string]("Error creating a user")
	}

	accessToken, err := u.AuthTokenService.GenerateToken(username, 1*time.Hour)
	if err != nil {
		return logic.FailedResult[string]("Error generating access token")
	}

	return logic.OkResult(accessToken)
}
