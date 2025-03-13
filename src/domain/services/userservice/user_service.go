package userservice

import "github.com/gabrielteiga/user-management-jwt/src/domain/interfaces"

type UserService struct {
	UserRepository interfaces.IUserRepository
}

func NewUserService(ur interfaces.IUserRepository) *UserService {
	return &UserService{
		UserRepository: ur,
	}
}

// TODO - implement the service that creates user and store in repository
