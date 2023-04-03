package repository

import (
	"example.com/user-service/src/application/logic"
	"example.com/user-service/src/domain/entity"
	"example.com/user-service/src/infra/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	Db *gorm.DB
}

func (r UserRepository) Create(user entity.User) *logic.Result[entity.User] {
	userModel := model.UserModel{
		Id:       user.Id,
		Username: user.Username,
		Password: user.Password,
	}
	userCreatedResult := r.Db.Create(&userModel)

	if userCreatedResult.Error != nil {
		return logic.FailedResult[entity.User]("Failed to create user")
	}

	createdUser := entity.User{Id: userModel.Id, Username: userModel.Username, Password: userModel.Password}
	return logic.OkResult(createdUser)
}

func (r UserRepository) GetByUsername(username string) *logic.Result[entity.User] {
	var existingUserModel model.UserModel

	existingUserResult := r.Db.Where("username = ?", username).First(&existingUserModel)

	if existingUserResult.Error != nil {
		return logic.FailedResult[entity.User]("Failed to get user")
	}

	existingUser := entity.User{Id: existingUserModel.Id, Username: existingUserModel.Username, Password: existingUserModel.Password}
	return logic.OkResult(existingUser)
}
