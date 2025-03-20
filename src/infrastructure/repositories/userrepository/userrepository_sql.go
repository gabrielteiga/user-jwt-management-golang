package userrepository

import (
	"database/sql"
	"log"

	"github.com/gabrielteiga/user-management-jwt/src/domain/entities"
)

type UserRepository struct {
	sql *sql.DB
}

func NewUserRepositorySQL(sql *sql.DB) *UserRepository {
	return &UserRepository{
		sql: sql,
	}
}

func (ur *UserRepository) Create(user *entities.User) error {
	sql := "INSERT INTO users (name, email, password) VALUES (?, ?, ?)"

	_, err := ur.sql.Exec(sql, user.Name, user.Email, user.Password)
	if err != nil {
		log.Printf("Error creating user: %v", err)
		return err
	}

	return nil
}
