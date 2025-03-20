package interfaces

import (
	"github.com/gabrielteiga/user-management-jwt/src/domain/entities"
)

type IUserRepository interface {
	Create(user *entities.User) error
}
