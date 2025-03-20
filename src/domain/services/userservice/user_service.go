package userservice

import (
	"github.com/gabrielteiga/user-management-jwt/src/domain/entities"
	"github.com/gabrielteiga/user-management-jwt/src/domain/interfaces"
)

type UserService struct {
	UserRepository interfaces.IUserRepository
}

func NewUserService(ur interfaces.IUserRepository) *UserService {
	return &UserService{
		UserRepository: ur,
	}
}

func (us *UserService) Create(user *entities.User) (*entities.User, error) {
	user.Role = entities.RoleUser

	err := us.UserRepository.Create(user)

	return user, err
}
