package UserRepository

import (
	"github.com/dev-newus/GoAlinBackend/src/models/entities"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) IUserRepository {
	return &Repository{
		db: db,
	}
}

func (repo *Repository) GetOneByEmail(email string) (*entities.User, error) {
	user := &entities.User{}
	result := repo.db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return user, result.Error
	}
	return user, nil
}

func (repo *Repository) Create(user entities.User) (entities.User, error) {
	hashedPass, err := hashPassword(user.Password)
	if err != nil {
		return user, err
	}

	user.Password = hashedPass

	result := repo.db.Create(&user)
	if result.Error != nil {
		return user, result.Error
	}

	return user, nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
