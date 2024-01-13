package UserRepository

import "github.com/dev-newus/GoAlinBackend/src/models/entities"

type IUserRepository interface {
	GetOneByEmail(email string) (*entities.User, error)
	Create(user entities.User) (entities.User, error)
}
