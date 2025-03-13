package cmd

import (
	"database/sql"
	"log"

	"github.com/gabrielteiga/user-management-jwt/src/api/controller/userscontrol"
	"github.com/gabrielteiga/user-management-jwt/src/api/routes"
	"github.com/gabrielteiga/user-management-jwt/src/domain/services/userservice"
	"github.com/gabrielteiga/user-management-jwt/src/infrastructure/repositories/userrepository"
	"github.com/gofiber/fiber/v2"
)

func Run() {
	app := fiber.New()

	db, _ := sql.Open("mysql", "root:toor@tcp(127.0.0.1:3306)/app_db")

	jr := routes.NewJobRouter(
		app,
		userscontrol.NewUserController(
			userservice.NewUserService(
				userrepository.NewUserRepositorySQL(db),
			),
		),
	)
	jr.CreateRoutes()

	log.Fatal(app.Listen(":3000"))
}
