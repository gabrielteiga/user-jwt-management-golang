package userscontrol

import "github.com/gabrielteiga/user-management-jwt/src/domain/services/userservice"

type UserController struct {
	UserService *userservice.UserService
}

func NewUserController(us *userservice.UserService) *UserController {
	return &UserController{
		UserService: us,
	}
}

// TODO - Create create method
