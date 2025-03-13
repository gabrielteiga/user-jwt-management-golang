package interfaces

import "github.com/gofiber/fiber/v2"

type IUserRepository interface {
	// TODO - remove c *fiber.Ctx and error return
	Create(c *fiber.Ctx) error
}
