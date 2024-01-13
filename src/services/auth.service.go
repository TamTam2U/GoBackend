package services

import (
	"github.com/dev-newus/GoAlinBackend/src/models/repositories/UserRepository"
)

type Repository struct {
	UserRepository.IUserRepository
}

func (repo *Repository) Register() {
	//logic of register
}
