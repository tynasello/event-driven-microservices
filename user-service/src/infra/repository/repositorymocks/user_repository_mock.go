package repositorymocks

import (
	"example.com/user-service/src/application/logic"
	"example.com/user-service/src/domain/entity"
	"gorm.io/gorm"
)

type UserRepositoryMock struct {
	Db    *gorm.DB
	Users []entity.User
}

func (rm *UserRepositoryMock) Create(user entity.User) *logic.Result[entity.User] {
	for _, item := range rm.Users {
		if item.Id == user.Id || item.Username == user.Username {
			return logic.FailedResult[entity.User]("Failed to create user")
		}
	}
	rm.Users = append(rm.Users, user)
	return logic.OkResult(user)
}

func (rm *UserRepositoryMock) GetByUsername(username string) *logic.Result[entity.User] {
	for _, item := range rm.Users {
		if item.Username == username {
			return logic.OkResult(item)
		}
	}
	return logic.FailedResult[entity.User]("Failed to get user")
}
