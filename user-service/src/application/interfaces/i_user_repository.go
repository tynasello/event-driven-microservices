package interfaces

import (
	"example.com/user-service/src/application/logic"
	"example.com/user-service/src/domain/entity"
)

type IUserRepository interface {
	Create(user entity.User) *logic.Result[entity.User]
	GetByUsername(username string) *logic.Result[entity.User]
}
