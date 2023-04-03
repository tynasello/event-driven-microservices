package usecase

import (
	"example.com/user-service/src/application/interfaces"
	"example.com/user-service/src/application/logic"
	"example.com/user-service/src/domain/entity"
)

type GetUserUseCase struct {
	UserRepository interfaces.IUserRepository
}

func (u GetUserUseCase) Execute(username string) *logic.Result[entity.User] {
	existingUserResult := u.UserRepository.GetByUsername(username)

	if existingUserResult.IsFailure {
		return logic.FailedResult[entity.User]("Failed to get user")
	}

	existingUser := existingUserResult.GetValue()

	return logic.OkResult(existingUser)
}
