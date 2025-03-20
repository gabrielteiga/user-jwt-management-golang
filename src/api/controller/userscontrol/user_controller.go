package userscontrol

import (
	"encoding/json"
	"log"

	"github.com/gabrielteiga/user-management-jwt/src/api/responses"
	"github.com/gabrielteiga/user-management-jwt/src/domain/entities"
	"github.com/gabrielteiga/user-management-jwt/src/domain/services/userservice"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	UserService *userservice.UserService
}

func NewUserController(us *userservice.UserService) *UserController {
	return &UserController{
		UserService: us,
	}
}

type CreateUserRequest struct {
	Name     string            `json:"name" validate:"required,min=3"`
	Email    string            `json:"email" validate:"required,email"`
	Password string            `json:"password,omitempty" validate:"required"`
	Role     entities.UserRole `json:"role,omitempty"`
}

func (uc *UserController) Create(c *fiber.Ctx) error {
	var user *entities.User
	json.Unmarshal(c.Request().Body(), &user)

	validate := validator.New()
	err := validate.Struct(&CreateUserRequest{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		Role:     user.Role,
	})
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.Error("Validation error", err.Error()))
	}

	log.Print("foi")
	log.Print(user)
	user, err = uc.UserService.Create(user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(responses.Error("Error creating user", err.Error()))
	}

	user.Password = ""
	return c.Status(fiber.StatusCreated).
		JSON(responses.Success("User created successfully", user))
}
