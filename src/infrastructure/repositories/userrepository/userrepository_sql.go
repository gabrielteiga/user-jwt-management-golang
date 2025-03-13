package userrepository

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
)

type UserRepository struct {
	sql *sql.DB
}

func NewUserRepositorySQL(sql *sql.DB) *UserRepository {
	return &UserRepository{
		sql: sql,
	}
}

func (ur *UserRepository) Create(c *fiber.Ctx) error {
	// TODO - Implement the functionality to create user in users table
	return c.SendString("OI MUNDO")
}
