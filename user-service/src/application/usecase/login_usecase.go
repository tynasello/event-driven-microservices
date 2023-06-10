package usecase

import (
	"time"

	"example.com/user-service/src/application/interfaces"
	"example.com/user-service/src/application/logic"
)

type LoginUseCase struct {
	UserRepository   interfaces.IUserRepository
	HashService      interfaces.IHashService
	AuthTokenService interfaces.IAuthTokenService
}

func (u LoginUseCase) Execute(username string, password string) *logic.Result[string] {
	existingUserResult := u.UserRepository.GetByUsername(username)

	if existingUserResult.IsFailure {
		return logic.FailedResult[string]("Failed to get user")
	}

	existingUser, _ := existingUserResult.GetValue()

	err := u.HashService.ValidateHash(existingUser.Password, password)
	if err != nil {
		return logic.FailedResult[string]("Invalid credentials")
	}

	accessToken, err := u.AuthTokenService.GenerateToken(username, 1*time.Hour)
	if err != nil {
		return logic.FailedResult[string]("Error generating access token")
	}

	return logic.OkResult(accessToken)
}
